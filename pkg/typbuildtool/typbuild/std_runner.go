package typbuild

import (
	"os"
	"os/exec"
)

// StdRunner is standard runner
type StdRunner struct {
}

// NewRunner return new instance of StdRunner
func NewRunner() *StdRunner {
	return &StdRunner{}
}

// Run the project
func (*StdRunner) Run(c *RunContext) (err error) {
	cmd := exec.CommandContext(c.Cli.Context, c.Binary, c.Cli.Args().Slice()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
