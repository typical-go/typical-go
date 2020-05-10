package typgo

import "github.com/urfave/cli/v2"

func cmdRun(c *Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFunc("RUN", run),
	}
}

func run(c *CliContext) (err error) {
	for _, module := range c.Core.BuildSequences {
		if runner, ok := module.(Runner2); ok {
			if err = runner.Run(c); err != nil {
				return
			}
		}
	}
	return
}
