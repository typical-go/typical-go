package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// SimpleCommander return command based on command function
type SimpleCommander struct {
	funcs []CommandFn
}

// CommandFn is a function to return command
type CommandFn func(ctx *Context) *cli.Command

// CreateCommander return new instance Commander
func CreateCommander(funcs ...CommandFn) *SimpleCommander {
	return &SimpleCommander{
		funcs: funcs,
	}
}

// Commands return list of command
func (c *SimpleCommander) Commands(ctx *Context) (cmds []*cli.Command) {
	for _, fn := range c.funcs {
		cmds = append(cmds, fn(ctx))
	}
	return
}
