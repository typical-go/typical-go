package typbuildtool

import "github.com/urfave/cli/v2"

func (b *BuildTool) cmdTest(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action: func(cliCtx *cli.Context) error {
			return b.Test(c.BuildContext(cliCtx))
		},
	}
}

// Test the project
func (b *BuildTool) Test(c *BuildContext) (err error) {
	for _, module := range b.buildSequences {
		if tester, ok := module.(Tester); ok {
			if err = tester.Test(c); err != nil {
				return
			}
		}
	}
	return
}
