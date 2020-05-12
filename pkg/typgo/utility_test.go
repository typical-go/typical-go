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

func TestUtilities(t *testing.T) {
	cmd1 := &cli.Command{}
	cmd2 := &cli.Command{}
	cmd3 := &cli.Command{}
	cmd4 := &cli.Command{}
	cmd5 := &cli.Command{}

	utilities := typgo.Utilities{
		typgo.NewUtility(func(*typgo.BuildTool) []*cli.Command {
			return []*cli.Command{cmd1}
		}),
		typgo.NewUtility(func(*typgo.BuildTool) []*cli.Command {
			return []*cli.Command{cmd2}
		}),
		typgo.Utilities{
			typgo.NewUtility(func(*typgo.BuildTool) []*cli.Command {
				return []*cli.Command{cmd3}
			}),
			typgo.NewUtility(func(*typgo.BuildTool) []*cli.Command {
				return []*cli.Command{cmd4, cmd5}
			}),
		},
	}

	require.Equal(t,
		[]*cli.Command{cmd1, cmd2, cmd3, cmd4, cmd5},
		utilities.Commands(nil),
	)
}
