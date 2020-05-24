package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func cmdClean(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project",
		Action:  c.ActionFn("CLEAN", clean),
	}
}

func clean(c *Context) (err error) {
	removeAll(c, typvar.BinFolder)

	build := typvar.GetBuild()
	remove(c, build.Binary)
	remove(c, build.Checksum)
	remove(c, build.Source)
	return
}

func removeAll(c *Context, folder string) {
	c.Infof("RemoveAll: %s", folder)
	if err := os.RemoveAll(folder); err != nil {
		c.Warn(err.Error())
	}
}

func remove(c *Context, file string) {
	c.Infof("Remove: %s", file)
	if err := os.Remove(file); err != nil {
		c.Warn(err.Error())
	}
}
