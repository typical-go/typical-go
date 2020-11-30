package typgo

import (
	"github.com/urfave/cli/v2"
)

type (
	// CliCommander interface return cli.Command method
	CliCommander interface {
		Cli(*BuildSys) *cli.Command
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

var _ CliCommander = (*Command)(nil)

// Cli command
func (c *Command) Cli(b *BuildSys) *cli.Command {
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
