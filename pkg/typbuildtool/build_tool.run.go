package typbuildtool

import "github.com/urfave/cli/v2"

func (b *BuildTool) cmdRun(c *Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action: func(cliCtx *cli.Context) (err error) {
			return b.Run(c.CliContext(cliCtx))
		},
	}
}

// Run task
func (b *BuildTool) Run(c *CliContext) (err error) {
	for _, module := range b.buildSequences {
		if runner, ok := module.(Runner); ok {
			if err = runner.Run(c); err != nil {
				return
			}
		}
	}
	return
}
