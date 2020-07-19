package typgo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCompileCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.CompileCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:    "compile",
				Aliases: []string{"c"},
				Usage:   "Compile the project",
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}

func TestStdCompile(t *testing.T) {
	cmpl := &typgo.StdCompile{}
	c := &typgo.Context{
		Descriptor: &typgo.Descriptor{Name: "some-name", Version: "0.0.1"},
		Context:    &cli.Context{Context: context.Background()},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: []string{
				"go", "build",
				"-ldflags", "-X github.com/typical-go/typical-go/pkg/typapp.Name=some-name -X github.com/typical-go/typical-go/pkg/typapp.Version=0.0.1",
				"-o", "bin/some-name",
				"./cmd/some-name",
			},
		},
	})
	defer unpatch(t)

	require.NoError(t, cmpl.Execute(c))
}

func TestStdCompile_OverrideParams(t *testing.T) {
	cmpl := &typgo.StdCompile{
		MainPackage: "some-package",
		Output:      "some-output",
		Ldflags: execkit.BuildVars{
			"some-var": "some-value",
		},
	}
	c := &typgo.Context{
		Descriptor: &typgo.Descriptor{Name: "some-name", Version: "0.0.1"},
		Context:    &cli.Context{Context: context.Background()},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: []string{
				"go", "build",
				"-ldflags", "-X some-var=some-value",
				"-o", "some-output",
				"some-package",
			},
		},
	})
	defer unpatch(t)

	require.NoError(t, cmpl.Execute(c))
}
