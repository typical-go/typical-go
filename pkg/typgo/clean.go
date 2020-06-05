package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// Cleaner responsible to clean
	Cleaner interface {
		Clean(*Context) error
	}
	// StdClean standard clean
	StdClean struct{}
)

//
// StdClean
//

var _ Cleaner = (*StdClean)(nil)

// Clean project
func (s *StdClean) Clean(c *Context) error {
	removeAll(c, typvar.BinFolder)
	removeAll(c, fmt.Sprintf("%s/bin", typvar.TypicalTmp))
	remove(c, typvar.BuildChecksum)
	removeAll(c, typvar.BuildToolSrc)
	remove(c, typvar.Precond(c.Descriptor.Name))
	return nil
}

func removeAll(c *Context, folder string) {
	if err := os.RemoveAll(folder); err == nil {
		c.Infof("RemoveAll: %s", folder)
	}
}

func remove(c *Context, file string) {
	if err := os.Remove(file); err == nil {
		c.Infof("Remove: %s", file)
	}
}

//
// commands
//

func cmdClean(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project",
		Action:  c.ActionFn("CLEAN", c.Clean.Clean),
	}
}
