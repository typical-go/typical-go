package typgo

import (
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// CleanCmd command clean
	CleanCmd struct {
		Action
	}
	// StdClean standard clean
	StdClean struct{}
)

//
// CleanCmd
//

var _ Cmd = (*CleanCmd)(nil)

// Command clean
func (c *CleanCmd) Command(b *BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "clean",
		Usage:  "Clean the project",
		Action: b.ActionFn("CLEAN", c.Execute),
	}
}

//
// StdClean
//

var _ Action = (*StdClean)(nil)

// Execute standard clean
func (s *StdClean) Execute(c *Context) error {
	removeAll(c, TypicalTmp)
	return nil
}

func removeAll(c *Context, folder string) {
	if err := os.RemoveAll(folder); err == nil {
		c.Infof("RemoveAll: %s", folder)
	}
}
