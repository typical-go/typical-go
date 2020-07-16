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
		Packages     []string
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
		Action:  b.ActionFn(t.Execute),
	}
}

//
// StdTest
//

var _ Action = (*StdTest)(nil)

// Execute standard test
func (s *StdTest) Execute(c *Context) (err error) {
	if len(c.Descriptor.Layouts) < 1 {
		fmt.Println("Nothing to test")
		return
	}

	return c.Execute(&execkit.GoTest{
		Packages:     s.GetPackages(c),
		Timeout:      s.GetTimeout(),
		CoverProfile: s.GetCoverProfile(),
		Race:         s.Race,
	})
}

// GetPackages return packages to be tested
func (s *StdTest) GetPackages(c *Context) []string {
	if len(s.Packages) < 1 {
		for _, layout := range c.Descriptor.Layouts {
			s.Packages = append(s.Packages, fmt.Sprintf("./%s/...", layout))
		}
	}
	return s.Packages
}

// GetTimeout return test timeout
func (s *StdTest) GetTimeout() time.Duration {
	if s.Timeout <= 0 {
		return defaultTestTimeout
	}
	return s.Timeout
}

// GetCoverProfile return cover profile filename
func (s *StdTest) GetCoverProfile() string {
	if s.CoverProfile == "" {
		return defaultTestCoverProfile
	}
	return s.CoverProfile
}
