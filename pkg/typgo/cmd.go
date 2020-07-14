package typgo

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

type (
	// Cmd interface return Command method
	Cmd interface {
		Command(*BuildCli) *cli.Command
	}
	// Commands list
	Commands []Cmd
	// Command to run action
	Command struct {
		Name            string
		Aliases         []string
		Flags           []cli.Flag
		SkipFlagParsing bool
		Usage           string
		Action          Action
	}
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
	// CmdTestCase test-case for cmd interface
	CmdTestCase struct {
		TestName      string
		Cmd           Cmd
		Expected      Command
		ExpectedError string
	}
)

//
// Command
//

var _ Cmd = (*Command)(nil)

// Command of command
func (c *Command) Command(b *BuildCli) *cli.Command {
	return &cli.Command{
		Name:            c.Name,
		Aliases:         c.Aliases,
		Flags:           c.Flags,
		SkipFlagParsing: c.SkipFlagParsing,

		Action: func(cliCtx *cli.Context) error {
			return c.Action.Execute(b.Context(cliCtx))
		},
	}
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
