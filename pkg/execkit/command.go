package execkit

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Command is wrapper to exec.Command
type Command struct {
	Name   string
	Args   []string
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
	Dir    string
	Env    []string
}

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

func (c Command) String() string {
	return fmt.Sprintf("%s %s", c.Name, strings.Join(c.Args, " "))
}

// Print command
func (c Command) Print(printer io.Writer) {
	color.New(color.FgMagenta).Fprint(printer, "\n$ ")
	fmt.Fprintf(printer, "%s ", c.Name)
	fmt.Fprintln(printer, strings.Join(c.Args, " "))
}
