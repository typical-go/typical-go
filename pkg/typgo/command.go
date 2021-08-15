package typgo

import (
	"context"
	"fmt"
	"io"
	"os"
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
	// Commander responsible to Command
	Commander interface {
		Command(extras ...string) *Command
	}
)

var _ Commander = (*Command)(nil)
var _ Action = (*Command)(nil)
var _ fmt.Stringer = (*Command)(nil)

// CommandLine create bash command line
func CommandLine(line string) *Command {
	slices := strings.Fields(line)
	return &Command{
		Name:   slices[0],
		Args:   slices[1:],
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
}

// ExecCmd return exec.Cmd
func (b *Command) ExecCmd(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, b.Name, b.Args...)
	cmd.Stdout = b.Stdout
	cmd.Stderr = b.Stderr
	cmd.Stdin = b.Stdin
	cmd.Dir = b.Dir
	cmd.Env = b.Env
	return cmd
}

// Command return Command
func (b *Command) Command(extras ...string) *Command {
	return b
}

// Execute bash
func (b *Command) Execute(c *Context) error {
	return c.ExecuteCommand(b)
}

func (b Command) String() string {
	var out strings.Builder
	fmt.Fprint(&out, b.Name)
	for _, arg := range b.Args {
		if strings.ContainsAny(arg, " ") {
			fmt.Fprintf(&out, " \"%s\"", arg)
		} else {
			fmt.Fprintf(&out, " %s", arg)
		}

	}
	return out.String()
}
