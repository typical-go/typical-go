package typgo

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/urfave/cli/v2"
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
	color.New(color.FgMagenta).Fprint(oskit.Stdout, "\n$ ")
	fmt.Fprintln(oskit.Stdout, cmd)
	return execkit.Run(c.Ctx(), cmd)
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}

//
// Actions
//

var _ Action = (Actions)(nil)

// Execute actions
func (a Actions) Execute(c *Context) error {
	for _, action := range a {
		if err := action.Execute(c); err != nil {
			return err
		}
	}
	return nil
}
