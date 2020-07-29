package typgo

import (
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
		Usage           string
		Flags           []cli.Flag
		SkipFlagParsing bool
		Action          Action
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
		Usage:           c.Usage,
		Flags:           c.Flags,
		SkipFlagParsing: c.SkipFlagParsing,

		Action: func(cliCtx *cli.Context) error {
			return b.Execute(c.Action, cliCtx)
		},
	}
}
