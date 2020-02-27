package typrun

import (
	"context"
	"os"
	"os/exec"
)

// StdRunner is standard runner
type StdRunner struct {
}

// New instance of StdRunner
func New() *StdRunner {
	return &StdRunner{}
}

// Run the project
func (*StdRunner) Run(ctx context.Context, c *Context) (err error) {
	cmd := exec.CommandContext(ctx, c.Binary, c.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
