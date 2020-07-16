package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestRunCompile(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.RunCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:            "run",
				Aliases:         []string{"r"},
				Usage:           "Run the project in local environment",
				SkipFlagParsing: true,
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}

func TestStdRun(t *testing.T) {
	t.Run("predefined", func(t *testing.T) {
		stdRun := &typgo.StdRun{
			Binary: "some-binary",
		}
		require.Equal(t, "some-binary", stdRun.GetBinary(nil))
	})
	t.Run("default", func(t *testing.T) {
		stdRun := &typgo.StdRun{}
		ctx := &typgo.Context{
			Descriptor: &typgo.Descriptor{Name: "some-name"},
		}
		require.Equal(t, "bin/some-name", stdRun.GetBinary(ctx))
	})
}
