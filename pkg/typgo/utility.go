package typgo

import (
	"github.com/urfave/cli/v2"
)

type (
	// Utility for build-tool
	Utility interface {
		Commands(*BuildCli) ([]*cli.Command, error)
	}
	// Utilities is list of utility
	Utilities []Utility
	// CommandFn is a function to return command
	CommandFn   func(ctx *BuildCli) ([]*cli.Command, error)
	utilityImpl struct {
		fn CommandFn
	}
)

//
// utilityImpl
//

var _ Utility = (*utilityImpl)(nil)

// NewUtility return new instance of utility
func NewUtility(fn CommandFn) Utility {
	return &utilityImpl{
		fn: fn,
	}
}

// CreateUtility to create utility from commands
func CreateUtility(cmds ...*cli.Command) Utility {
	return NewUtility(func(ctx *BuildCli) ([]*cli.Command, error) {
		return cmds, nil
	})
}

// Commands of SimpleUtility
func (s *utilityImpl) Commands(b *BuildCli) ([]*cli.Command, error) {
	return s.fn(b)
}

//
// Utilities
//

var _ Utility = (Utilities)(nil)

// Commands of Utilities
func (u Utilities) Commands(b *BuildCli) ([]*cli.Command, error) {
	var cmds []*cli.Command
	for _, utility := range u {
		cmds0, err := utility.Commands(b)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmds0...)
	}
	return cmds, nil
}
