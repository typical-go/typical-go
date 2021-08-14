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
	// Bash is wrapper to exec.Bash
	Bash struct {
		Name   string
		Args   []string
		Stdout io.Writer
		Stderr io.Writer
		Stdin  io.Reader
		Dir    string
		Env    []string
	}
	// Basher responsible to Bash
	Basher interface {
		Bash(extras ...string) *Bash
	}
)

var _ Basher = (*Bash)(nil)
var _ Action = (*Bash)(nil)
var _ fmt.Stringer = (*Bash)(nil)

// BashCommand create bash command line
func BashCommand(line string) *Bash {
	slices := strings.Split(line, " ")
	return &Bash{
		Name:   slices[0],
		Args:   slices[1:],
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
}

// ExecCmd return exec.Cmd
func (b *Bash) ExecCmd(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, b.Name, b.Args...)
	cmd.Stdout = b.Stdout
	cmd.Stderr = b.Stderr
	cmd.Stdin = b.Stdin
	cmd.Dir = b.Dir
	cmd.Env = b.Env
	return cmd
}

// Bash return Bash
func (b *Bash) Bash(extras ...string) *Bash {
	return b
}

// Execute bash
func (b *Bash) Execute(c *Context) error {
	return c.Execute(b)
}

func (b Bash) String() string {
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
