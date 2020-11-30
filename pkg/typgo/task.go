package typgo

import (
	"github.com/urfave/cli/v2"
)

type (
	// Tasker interface return cli.Command method
	Tasker interface {
		Task(*BuildSys) *cli.Command
	}
	// Task to run action
	Task struct {
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

var _ Tasker = (*Task)(nil)

// Task command
func (c *Task) Task(b *BuildSys) *cli.Command {
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
