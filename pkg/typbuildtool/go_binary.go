package typbuildtool

import (
	"io"
	"os/exec"
)

// GoBinary is go binary istribution
type GoBinary struct {
	binary string
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

// Run the build distribution
func (d *GoBinary) Run(c *BuildContext) (err error) {
	cmd := exec.CommandContext(c.Cli.Context, d.binary, c.Cli.Args().Slice()...)
	cmd.Stdout = d.stdout
	cmd.Stderr = d.stderr
	cmd.Stdin = d.stdin
	return cmd.Run()
}
