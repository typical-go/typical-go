package typgo

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

type (
	// Cmd interface return Command method
	Cmd interface {
		Command(*BuildSys) *cli.Command
	}
	// Command to run action
	Command struct {
		Name            string
		Aliases         []string
		Flags           []cli.Flag
		SkipFlagParsing bool
		Usage           string
		Action          Action
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
func (c *Command) Command(b *BuildSys) *cli.Command {
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
// CmdTestCase
//

// Run test for Command
func (tt *CmdTestCase) Run(t *testing.T) bool {
	return t.Run(tt.TestName, func(t *testing.T) {
		b := &BuildSys{}
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
