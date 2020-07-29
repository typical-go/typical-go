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
		Name    string   // By default is "test"
		Aliases []string // By default is "t"
		Usage   string   // By default is "Test the project"
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
func (t *TestCmd) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    t.getName(),
		Aliases: t.getAliases(),
		Usage:   t.getUsage(),
		Action:  b.ActionFn(t.Action),
	}
}

func (t *TestCmd) getName() string {
	if t.Name == "" {
		t.Name = "test"
	}
	return t.Name
}

func (t *TestCmd) getAliases() []string {
	if len(t.Aliases) < 1 {
		t.Aliases = []string{"t"}
	}
	return t.Aliases
}

func (t *TestCmd) getUsage() string {
	if t.Usage == "" {
		t.Usage = "Test the project"
	}
	return t.Usage
}

//
// StdTest
//

var _ Action = (*StdTest)(nil)

// Execute standard test
func (s *StdTest) Execute(c *Context) (err error) {
	if len(c.BuildSys.Layouts) < 1 {
		fmt.Println("Nothing to test")
		return
	}

	return c.Execute(&execkit.GoTest{
		Packages:     s.getPackages(c),
		Timeout:      s.getTimeout(),
		CoverProfile: s.getCoverProfile(),
		Race:         s.Race,
	})
}

func (s *StdTest) getPackages(c *Context) []string {
	if len(s.Packages) < 1 {
		for _, layout := range c.BuildSys.Layouts {
			s.Packages = append(s.Packages, fmt.Sprintf("./%s/...", layout))
		}
	}
	return s.Packages
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
