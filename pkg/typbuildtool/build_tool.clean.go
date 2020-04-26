package typbuildtool

import (
	"os"

	"github.com/urfave/cli/v2"
)

func (b *BuildTool) cmdClean(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project",
		Action: func(cliCtx *cli.Context) (err error) {
			return b.Clean(c.CliContext(cliCtx))
		},
	}
}

// Clean the project
func (b *BuildTool) Clean(c *CliContext) (err error) {
	for _, module := range b.buildSequences {
		if cleaner, ok := module.(Cleaner); ok {
			if err = cleaner.Clean(c); err != nil {
				return
			}
		}
	}

	c.Infof("Remove All: %s", c.TypicalTmp)
	if err := os.RemoveAll(c.TypicalTmp); err != nil {
		c.Warn(err.Error())
	}

	return
}
