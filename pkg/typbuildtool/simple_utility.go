package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// SimpleUtility return command based on command function
type SimpleUtility struct {
	fn UtilityFn
}

// UtilityFn is a function to return command
type UtilityFn func(ctx *Context) []*cli.Command

// NewUtility return new instance SimpleUtility
func NewUtility(fn UtilityFn) *SimpleUtility {
	return &SimpleUtility{
		fn: fn,
	}
}

// Commands return list of command
func (c *SimpleUtility) Commands(ctx *Context) (cmds []*cli.Command) {
	return c.fn(ctx)
}
