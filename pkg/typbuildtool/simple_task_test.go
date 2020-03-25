package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"

	"github.com/urfave/cli/v2"
)

func TestSimpleCommander(t *testing.T) {
	t.Run("SHOULD implement Commanders", func(t *testing.T) {
		var _ typbuildtool.Task = typbuildtool.NewTask()
	})
}

func TestSimpleCommander_Commands(t *testing.T) {
	cmd1 := &cli.Command{}
	cmd2 := &cli.Command{}

	commander := typbuildtool.NewTask(
		func(ctx *typbuildtool.Context) *cli.Command {
			return cmd1
		},
		func(ctx *typbuildtool.Context) *cli.Command {
			return cmd2
		},
	)

	require.Equal(t, []*cli.Command{cmd1, cmd2}, commander.Commands(nil))
}
