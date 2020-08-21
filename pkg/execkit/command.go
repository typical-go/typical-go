package execkit

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type (
	// Command is wrapper to exec.Command
	Command struct {
		Name   string
		Args   []string
		Stdout io.Writer
		Stderr io.Writer
		Stdin  io.Reader
		Dir    string
		Env    []string
	}
	// Commander responsible to command
	Commander interface {
		Command() *Command
	}
)

var _ Commander = (*Command)(nil)
var _ fmt.Stringer = (*Command)(nil)

// Run the comand
func (c *Command) Run(ctx context.Context) (err error) {
	return c.ExecCmd(ctx).Run()
}

// ExecCmd return exec.Cmd
func (c *Command) ExecCmd(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.Name, c.Args...)
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stderr
	cmd.Stdin = c.Stdin
	cmd.Dir = c.Dir
	cmd.Env = c.Env
	return cmd
}

// Command return command
func (c *Command) Command() *Command {
	return c
}

func (c Command) String() string {
	var out strings.Builder
	fmt.Fprint(&out, c.Name)
	for _, arg := range c.Args {
		if strings.ContainsAny(arg, " ") {
			fmt.Fprintf(&out, " \"%s\"", arg)
		} else {
			fmt.Fprintf(&out, " %s", arg)
		}

	}
	return out.String()
}
