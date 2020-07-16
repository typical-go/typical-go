package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCleanCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.CleanCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:  "clean",
				Usage: "Clean the project",
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}

func TestStdClean(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		typgo.TypicalTmp = ".typical-tmp"
		stdClean := &typgo.StdClean{}
		require.Equal(t, []string{".typical-tmp"}, stdClean.GetPaths())
	})
	t.Run("predefined", func(t *testing.T) {
		stdClean := &typgo.StdClean{
			Paths: []string{"path1", "path2"},
		}
		require.Equal(t, []string{"path1", "path2"}, stdClean.GetPaths())
	})
}
