package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCleanCmd(t *testing.T) {
	cleanCmd := &typgo.CleanCmd{
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := cleanCmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")
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
