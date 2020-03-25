package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// SimpleTask return command based on command function
type SimpleTask struct {
	funcs []CommandFn
}

// CommandFn is a function to return command
type CommandFn func(ctx *Context) *cli.Command

// NewTask return new instance Commander
func NewTask(funcs ...CommandFn) *SimpleTask {
	return &SimpleTask{
		funcs: funcs,
	}
}

// Commands return list of command
func (c *SimpleTask) Commands(ctx *Context) (cmds []*cli.Command) {
	for _, fn := range c.funcs {
		cmds = append(cmds, fn(ctx))
	}
	return
}
