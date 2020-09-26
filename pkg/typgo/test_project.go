package typgo

import (
	"fmt"
	"time"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

const (
	defaultTestTimeout      = 25 * time.Second
	defaultTestCoverProfile = "cover.out"
)

type (
	// TestProject command test
	TestProject struct {
		Timeout      time.Duration
		CoverProfile string
		Race         bool
		Packages     []string
	}
)

var _ Cmd = (*TestProject)(nil)
var _ Action = (*TestProject)(nil)

// Command test
func (t *TestProject) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  b.Action(t),
	}
}

// Execute standard test
func (t *TestProject) Execute(c *Context) (err error) {
	if len(c.BuildSys.ProjectLayouts) < 1 {
		fmt.Println("Nothing to test")
		return
	}

	return c.Execute(&execkit.GoTest{
		Packages:     t.getPackages(c),
		Timeout:      t.getTimeout(),
		CoverProfile: t.getCoverProfile(),
		Race:         t.Race,
	})
}

func (t *TestProject) getPackages(c *Context) []string {
	if len(t.Packages) < 1 {
		for _, layout := range c.BuildSys.ProjectLayouts {
			t.Packages = append(t.Packages, fmt.Sprintf("./%s/...", layout))
		}
	}
	return t.Packages
}

func (t *TestProject) getTimeout() time.Duration {
	if t.Timeout <= 0 {
		return defaultTestTimeout
	}
	return t.Timeout
}

func (t *TestProject) getCoverProfile() string {
	if t.CoverProfile == "" {
		return defaultTestCoverProfile
	}
	return t.CoverProfile
}
