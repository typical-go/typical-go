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
		Action:  c.ActionFunc(b.Clean),
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

	typicalTmp := c.Core.TypicalTmp

	c.Infof("Remove All: %s", typicalTmp)
	if err := os.RemoveAll(typicalTmp); err != nil {
		c.Warn(err.Error())
	}

	return
}
