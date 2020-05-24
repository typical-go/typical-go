package typgo

import "github.com/urfave/cli/v2"

func cmdTest(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  c.ActionFn("TEST", test),
	}
}

func test(c *Context) (err error) {
	_, err = c.Execute(c, TestPhase)
	return
}
