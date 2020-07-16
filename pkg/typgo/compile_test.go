package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
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
	t.Run("predefine", func(t *testing.T) {

		cmpl := &typgo.StdCompile{
			MainPackage: "some-package",
			Output:      "some-output",
			Ldflags: execkit.BuildVars{
				"some-var": "some-value",
			},
		}
		require.Equal(t, "-X some-var=some-value", cmpl.GetLdflags(nil).String())
		require.Equal(t, "some-package", cmpl.GetMainPackage(nil))
		require.Equal(t, "some-output", cmpl.GetOutput(nil))

	})
	t.Run("default", func(t *testing.T) {
		cmpl := &typgo.StdCompile{}
		ctx := &typgo.Context{
			Descriptor: &typgo.Descriptor{
				Name:    "some-name",
				Version: "0.0.1",
			},
		}

		require.Equal(t, "-X github.com/typical-go/typical-go/pkg/typapp.Name=some-name -X github.com/typical-go/typical-go/pkg/typapp.Version=0.0.1", cmpl.GetLdflags(ctx).String())
		require.Equal(t, "./cmd/some-name", cmpl.GetMainPackage(ctx))
		require.Equal(t, "bin/some-name", cmpl.GetOutput(ctx))
	})
}
