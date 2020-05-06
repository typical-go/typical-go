package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

var (
	_ Utility = (*SimpleUtility)(nil)
)

type (
	// Utility of build-tool
	Utility interface {
		Commands(c *Context) []*cli.Command
	}

	// SimpleUtility return command based on command function
	SimpleUtility struct {
		fn UtilityFn
	}

	// UtilityFn is a function to return command
	UtilityFn func(ctx *Context) []*cli.Command
)

// NewUtility return new instance of utility
func NewUtility(fn UtilityFn) *SimpleUtility {
	return &SimpleUtility{
		fn: fn,
	}
}

// Commands return list of command
func (s *SimpleUtility) Commands(ctx *Context) (cmds []*cli.Command) {
	return s.fn(ctx)
}
