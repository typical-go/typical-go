package typgo

import (
	"github.com/urfave/cli/v2"
)

var (
	_ Utility = (*SimpleUtility)(nil)
	_ Utility = (Utilities)(nil)
)

type (
	// Utility for build-tool
	Utility interface {
		Commands(*BuildTool) []*cli.Command
	}

	// Utilities is list of utility
	Utilities []Utility

	// SimpleUtility return command based on command function
	SimpleUtility struct {
		fn UtilityFn
	}

	// UtilityFn is a function to return command
	UtilityFn func(ctx *BuildTool) []*cli.Command
)

// NewUtility return new instance of utility
func NewUtility(fn UtilityFn) *SimpleUtility {
	return &SimpleUtility{
		fn: fn,
	}
}

// Commands of SimpleUtility
func (s *SimpleUtility) Commands(b *BuildTool) []*cli.Command {
	return s.fn(b)
}

// Commands of Utilities
func (u Utilities) Commands(b *BuildTool) (cmds []*cli.Command) {
	for _, utility := range u {
		cmds = append(cmds, utility.Commands(b)...)
	}
	return
}
