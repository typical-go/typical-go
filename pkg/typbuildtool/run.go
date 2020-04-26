package typbuildtool

import "github.com/urfave/cli/v2"

func cmdRun(c *Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFunc(run),
	}
}

func run(c *CliContext) (err error) {
	for _, module := range c.BuildTool.buildSequences {
		if runner, ok := module.(Runner); ok {
			if err = runner.Run(c); err != nil {
				return
			}
		}
	}
	return
}
