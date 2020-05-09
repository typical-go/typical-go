package typbuild

import (
	"github.com/urfave/cli/v2"
)

var (
	_ Utility = (*SimpleUtility)(nil)
)

type (
	// Utility for build-tool
	Utility interface {
		Commands(c *Context) []*cli.Command
	}

	// Utilities is list of utility
	Utilities []Utility

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

// Commands of SimpleUtility
func (s *SimpleUtility) Commands(ctx *Context) []*cli.Command {
	return s.fn(ctx)
}

// Commands of Utilities
func (s *Utilities) Commands(ctx *Context) (cmds []*cli.Command) {
	return
}
