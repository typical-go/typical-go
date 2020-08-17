package typgo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// CleanProject command clean
	CleanProject struct {
		Paths []string
	}
)

var _ Cmd = (*CleanProject)(nil)
var _ Action = (*CleanProject)(nil)

// Command clean
func (p *CleanProject) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:   "clean",
		Usage:  "Clean the project",
		Action: b.Action(p),
	}
}

// Execute standard clean
func (p *CleanProject) Execute(c *Context) error {
	for _, path := range p.getPaths() {
		fmt.Fprintf(Stdout, "Removing %s\n", path)
		os.RemoveAll(path)
	}
	return nil
}

func (p *CleanProject) getPaths() []string {
	if len(p.Paths) < 1 {
		p.Paths = []string{TypicalTmp}
	}
	return p.Paths
}
