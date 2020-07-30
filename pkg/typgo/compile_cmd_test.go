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
	compileCmd := &typgo.CompileCmd{}
	command := compileCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "compile", command.Name)
	require.Equal(t, []string{"c"}, command.Aliases)
	require.Equal(t, "Compile the project", command.Usage)
	require.NoError(t, command.Action(&cli.Context{}))
}

func TestCompileCmd_Define(t *testing.T) {
	compileCmd := &typgo.CompileCmd{
		Name:    "some-name",
		Usage:   "some-usage",
		Aliases: []string{"x"},
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("action-error")
		}),
	}
	command := compileCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "some-name", command.Name)
	require.Equal(t, []string{"x"}, command.Aliases)
	require.Equal(t, "some-usage", command.Usage)
	require.EqualError(t, command.Action(&cli.Context{}), "action-error")
}

func TestStdCompile(t *testing.T) {
	cmpl := &typgo.StdCompile{}
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name", ProjectVersion: "0.0.1"},
		},
		Context: &cli.Context{Context: context.Background()},
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

func TestStdCompile_Predefined(t *testing.T) {
	cmpl := &typgo.StdCompile{
		MainPackage: "some-package",
		Output:      "some-output",
		Ldflags: execkit.BuildVars{
			"some-var": "some-value",
		},
	}
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name", ProjectVersion: "0.0.1"},
		},
		Context: &cli.Context{Context: context.Background()},
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
