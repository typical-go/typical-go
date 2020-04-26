package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// Test the project
func (b *StdBuild) Test(c *CliContext) (err error) {
	c.Info("Standard-Build: Test the project")
	var targets []string
	for _, source := range c.Core.AppSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	gotest := buildkit.NewGoTest(targets...).
		WithCoverProfile(b.coverProfile).
		WithTimeout(b.testTimeout).
		WithRace(true).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	return gotest.Execute(c.Context)
}
