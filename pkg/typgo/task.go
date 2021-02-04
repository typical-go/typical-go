package typgo

import (
	"github.com/urfave/cli/v2"
)

type (
	// Tasker interface return cli.Command method
	Tasker interface {
		Task() *Task
	}
	// Task to run action
	Task struct {
		Name            string
		Aliases         []string
		Usage           string
		Flags           []cli.Flag
		SkipFlagParsing bool
		Action          Action
		Before          Action
		SubTasks        []*Task
	}
)

//
// Command
//

var _ Tasker = (*Task)(nil)

// Task command
func (t *Task) Task() *Task {
	return t
}

// CliCommand return cli command from task
func CliCommand(d *Descriptor, t *Task) *cli.Command {
	cmd := &cli.Command{
		Name:            t.Name,
		Aliases:         t.Aliases,
		Usage:           t.Usage,
		Flags:           t.Flags,
		SkipFlagParsing: t.SkipFlagParsing,
		Action:          CliFunc(d, t.Action),
		Before:          CliFunc(d, t.Before),
	}
	for _, subTask := range t.SubTasks {
		cmd.Subcommands = append(cmd.Subcommands, CliCommand(d, subTask))
	}
	return cmd
}

// CliFunc return urfave cli function from Action
func CliFunc(d *Descriptor, a Action) func(*cli.Context) error {
	return func(c *cli.Context) error {
		if a == nil {
			return nil
		}

		return a.Execute(&Context{
			Context:    c,
			Descriptor: d,
		})
	}
}
