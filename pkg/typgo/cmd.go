package typgo

import (
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
