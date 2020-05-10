package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func cmdClean(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project",
		Action:  c.ActionFunc("CLEAN", clean),
	}
}

func clean(c *CliContext) (err error) {
	for _, module := range c.Core.BuildSequences {
		if cleaner, ok := module.(Cleaner); ok {
			if err = cleaner.Clean(c); err != nil {
				return
			}
		}
	}

	typicalTmp := typvar.TypicalTmp

	c.Infof("Remove All: %s", typicalTmp)
	if err := os.RemoveAll(typicalTmp); err != nil {
		c.Warn(err.Error())
	}

	return
}
