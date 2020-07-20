package typgo

import (
	"context"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

var (
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)

type (
	// Action responsible to execute process
	Action interface {
		Execute(*Context) error
	}
	// Actions for composite execution
	Actions []Action
	// ExecuteFn is execution function
	ExecuteFn  func(*Context) error
	actionImpl struct {
		fn ExecuteFn
	}
	// Context of build tool
	Context struct {
		*cli.Context
		BuildSys *BuildSys
	}
)

//
// actionImpl
//

// NewAction return new instance of Action
func NewAction(fn ExecuteFn) Action {
	return &actionImpl{
		fn: fn,
	}
}

// Execute action
func (a *actionImpl) Execute(c *Context) error {
	return a.fn(c)
}

//
// Context
//

// Execute command
func (c *Context) Execute(cmder execkit.Commander) error {
	cmd := cmder.Command()
	cmd.Print(Stdout)
	return execkit.Run(c.Ctx(), cmd)
}

// func (c *Context) Run() {
// 	for _, cmd := range c.Descriptor.Cmds {
// 		cmd.Command()
// 	}
// }

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}
