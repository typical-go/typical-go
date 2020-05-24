package typgo

import "github.com/urfave/cli/v2"

func cmdRun(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFn("RUN", run),
	}
}

func run(c *Context) (err error) {
	_, err = c.Execute(c, RunPhase)
	return
}
