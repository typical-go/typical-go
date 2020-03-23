package buildkit

import (
	"context"
	"io"
	"os/exec"
)

// Command of bash
type Command struct {
	name   string
	args   []string
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
	dir    string
}

// NewCommand return new instance of Command
func NewCommand(name string, args ...string) *Command {
	return &Command{
		name: name,
		args: args,
	}
}

// WithStdout return Command with new stdout
func (c *Command) WithStdout(stdout io.Writer) *Command {
	c.stdout = stdout
	return c
}

// WithStderr return Command with new stderr
func (c *Command) WithStderr(stderr io.Writer) *Command {
	c.stderr = stderr
	return c
}

// WithStdin return Command with new stdin
func (c *Command) WithStdin(stdin io.Reader) *Command {
	c.stdin = stdin
	return c
}

// WithDir return Command with new dir
func (c *Command) WithDir(dir string) *Command {
	c.dir = dir
	return c
}

// Execute comand
func (c *Command) Execute(ctx context.Context) (err error) {
	cmd := exec.CommandContext(ctx, c.name, c.Args()...)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr
	cmd.Stdin = c.stdin
	cmd.Dir = c.dir
	return cmd.Run()
}

// Args of command
func (c *Command) Args() []string {
	return c.args
}
