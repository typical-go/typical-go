package typgo

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

type (
	// BuildCmdRuns run command in current BuildSys
	BuildCmdRuns []string
)

var _ Action = (BuildCmdRuns)(nil)

// Execute BuildSysRuns
func (r BuildCmdRuns) Execute(c *Context) error {
	for _, name := range r {
		if err := RunCommand(c, name); err != nil {
			return err
		}
	}
	return nil
}

// RunCommand run command by name
func RunCommand(c *Context, name string) error {
	return runCommand(c.Context, c.BuildSys.Commands, strings.Split(name, "."), 0)
}

func runCommand(c *cli.Context, commands []*cli.Command, names []string, i int) error {
	for _, command := range commands {
		if command.Name == names[i] {
			if len(names) > i+1 {
				return runCommand(c, command.Subcommands, names, i+1)
			}
			return command.Action(c)
		}
	}
	return fmt.Errorf("typgo: %s not found", strings.Join(names, "."))
}
