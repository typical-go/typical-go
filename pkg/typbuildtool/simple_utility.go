package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// SimpleUtility return command based on command function
type SimpleUtility struct {
	funcs []UtilityFn
}

// UtilityFn is a function to return command
type UtilityFn func(ctx *Context) *cli.Command

// NewUtility return new instance SimpleUtility
func NewUtility(funcs ...UtilityFn) *SimpleUtility {
	return &SimpleUtility{
		funcs: funcs,
	}
}

// Commands return list of command
func (c *SimpleUtility) Commands(ctx *Context) (cmds []*cli.Command) {
	for _, fn := range c.funcs {
		cmds = append(cmds, fn(ctx))
	}
	return
}
