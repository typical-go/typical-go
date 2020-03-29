package typapp

import "github.com/urfave/cli/v2"

var (
	_ Commander = (*SimpleCommander)(nil)
)

// SimpleCommander is simple implementation of Commander
type SimpleCommander struct {
	fns []CommandFn
}

// CommandFn is a function to return command
type CommandFn func(ctx *Context) []*cli.Command

// NewCommander return new instance of SimpleComander
func NewCommander(fns ...CommandFn) *SimpleCommander {
	return &SimpleCommander{
		fns: fns,
	}
}

// Commands of application
func (s *SimpleCommander) Commands(c *Context) (cmds []*cli.Command) {
	for _, fn := range s.fns {
		cmds = append(cmds, fn(c)...)
	}
	return
}
