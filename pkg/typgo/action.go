package typgo

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typast"
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
		*Descriptor
		ASTStore *typast.ASTStore
		Imports  []string
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
// CmdTestCase
//

// Run test for Command
func (tt *CmdTestCase) Run(t *testing.T) bool {
	return t.Run(tt.TestName, func(t *testing.T) {
		b := &BuildCli{}
		cmd := tt.Cmd.Command(b)
		require.Equal(t, tt.Expected.Name, cmd.Name)
		require.Equal(t, tt.Expected.Usage, cmd.Usage)
		require.Equal(t, tt.Expected.Aliases, cmd.Aliases)
		require.Equal(t, tt.Expected.Flags, cmd.Flags)

		err := cmd.Action(nil)
		if tt.ExpectedError != "" {
			require.EqualError(t, err, tt.ExpectedError)
		} else {
			require.NoError(t, err)
		}
	})
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

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}
