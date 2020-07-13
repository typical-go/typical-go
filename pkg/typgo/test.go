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
	// TestCmd command test
	TestCmd struct {
		Action
	}
	// StdTest standard test
	StdTest struct {
		Timeout      time.Duration
		CoverProfile string
		Race         bool
	}
)

//
// TestCmd
//

// Command test
func (t *TestCmd) Command(b *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  b.ActionFn("TEST", t.Execute),
	}
}

//
// StdTest
//

var _ Action = (*StdTest)(nil)

// Execute standard test
func (s *StdTest) Execute(c *Context) (err error) {
	if len(c.Descriptor.Layouts) < 1 {
		c.Info("Nothing to test")
		return
	}

	return c.Execute(&execkit.GoTest{
		Targets:      testTargets(c),
		Timeout:      s.getTimeout(),
		CoverProfile: s.getCoverProfile(),
		Race:         s.Race,
	})
}

func testTargets(c *Context) (targets []string) {
	for _, layout := range c.Descriptor.Layouts {
		targets = append(targets, fmt.Sprintf("./%s/...", layout))
	}
	return
}

func (s *StdTest) getTimeout() time.Duration {
	if s.Timeout <= 0 {
		return defaultTestTimeout
	}
	return s.Timeout
}

func (s *StdTest) getCoverProfile() string {
	if s.CoverProfile == "" {
		return defaultTestCoverProfile
	}
	return s.CoverProfile
}
