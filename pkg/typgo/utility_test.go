package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"

	"github.com/urfave/cli/v2"
)

func TestSimpleCommander_Commands(t *testing.T) {
	expected := []*cli.Command{&cli.Command{}, &cli.Command{}, &cli.Command{}}
	expectedErr := errors.New("")

	utility := typgo.NewUtility(func(ctx *typgo.BuildCli) ([]*cli.Command, error) {
		return expected, expectedErr
	})

	cmds, err := utility.Commands(nil)
	require.Equal(t, expected, cmds)
	require.Equal(t, expectedErr, err)
}

func TestUtilities(t *testing.T) {
	cmd1 := &cli.Command{}
	cmd2 := &cli.Command{}
	cmd3 := &cli.Command{}
	cmd4 := &cli.Command{}
	cmd5 := &cli.Command{}

	utilities := typgo.Utilities{
		typgo.CreateUtility(cmd1),
		typgo.CreateUtility(cmd2),
		typgo.Utilities{
			typgo.CreateUtility(cmd3),
			typgo.CreateUtility(cmd4, cmd5),
		},
	}

	cmds, err := utilities.Commands(nil)

	require.Equal(t, []*cli.Command{cmd1, cmd2, cmd3, cmd4, cmd5}, cmds)
	require.NoError(t, err)
}
