package typgo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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
		Descriptor *Descriptor
	}
)

//
// actionImpl
//

// NewAction return new instance of Action
func NewAction(fn ExecuteFn) Action {
	return &actionImpl{fn: fn}
}

// Execute action
func (a *actionImpl) Execute(c *Context) error {
	return a.fn(c)
}

//
// Context
//

// Execute command
func (c *Context) Execute(basher Basher) error {
	bash := basher.Bash()
	color.New(color.FgMagenta).Fprint(oskit.Stdout, "$ ")
	fmt.Fprintln(oskit.Stdout, bash)
	return RunBash(c.Ctx(), bash)
}

// ExecuteBash execute bash command
func (c *Context) ExecuteBash(commandLine string) error {
	if commandLine == "" {
		return errors.New("command line can't be empty")
	}
	slices := strings.Split(commandLine, " ")
	return c.Execute(&Bash{
		Name:   slices[0],
		Args:   slices[1:],
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
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
