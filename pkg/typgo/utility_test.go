package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"

	"github.com/urfave/cli/v2"
)

func TestSimpleCommander_Commands(t *testing.T) {
	cmd1 := &cli.Command{}
	cmd2 := &cli.Command{}
	cmd3 := &cli.Command{}

	utility := typgo.NewUtility(func(ctx *typgo.BuildTool) []*cli.Command {
		return []*cli.Command{cmd1, cmd2, cmd3}
	})

	require.Equal(t, []*cli.Command{cmd1, cmd2, cmd3}, utility.Commands(nil))
}
