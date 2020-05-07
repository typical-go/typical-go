package typbuild

import "github.com/urfave/cli/v2"

func cmdTest(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  c.ActionFunc("TEST", test),
	}
}

func test(c *CliContext) (err error) {
	for _, module := range c.BuildTool.BuildSequences {
		if tester, ok := module.(Tester); ok {
			if err = tester.Test(c); err != nil {
				return
			}
		}
	}
	return
}
