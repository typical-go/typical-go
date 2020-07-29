package typgo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// CleanCmd command clean
	CleanCmd struct {
		Name  string // By default is "clean"
		Usage string // By default is "Clean the project"
		Action
	}
	// StdClean standard clean
	StdClean struct {
		Paths []string
	}
)

//
// CleanCmd
//

var _ Cmd = (*CleanCmd)(nil)

// Command clean
func (c *CleanCmd) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:   c.getName(),
		Usage:  c.getUsage(),
		Action: b.ActionFn(c.Action),
	}
}

func (c *CleanCmd) getName() string {
	if c.Name == "" {
		c.Name = "clean"
	}
	return c.Name
}

func (c *CleanCmd) getUsage() string {
	if c.Usage == "" {
		c.Usage = "Clean the project"
	}
	return c.Usage
}

//
// StdClean
//

var _ Action = (*StdClean)(nil)

// Execute standard clean
func (s *StdClean) Execute(c *Context) error {
	for _, path := range s.GetPaths() {
		if err := os.RemoveAll(path); err != nil {
			fmt.Printf("Failed removing %s\n", path)
		} else {
			fmt.Printf("Removing %s\n", path)
		}

	}
	// removeAll(c, TypicalTmp)
	return nil
}

// GetPaths return paths to be clean
func (s *StdClean) GetPaths() []string {
	if len(s.Paths) < 1 {
		s.Paths = []string{TypicalTmp}
	}
	return s.Paths
}
