package typbuildtool

import (
	"os"
	"os/exec"
)

// StdBuildDistribution is standard build distribution
type StdBuildDistribution struct {
	binary string
}

// NewBuildDistribution return new instance of BuildDistribution
func NewBuildDistribution(binary string) *StdBuildDistribution {
	return &StdBuildDistribution{
		binary: binary,
	}
}

// Run the build distribution
func (d *StdBuildDistribution) Run(c *Context) (err error) {
	cmd := exec.CommandContext(c.Cli.Context, d.binary, c.Cli.Args().Slice()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
