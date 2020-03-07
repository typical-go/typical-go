package typbuild

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"

	log "github.com/sirupsen/logrus"
)

// StdTester is standard tester
type StdTester struct {
	coverProfile string
}

// NewTester return new instance of StdTester
func NewTester() *StdTester {
	return &StdTester{
		coverProfile: "cover.out",
	}
}

// WithCoverProfile return StdTester with new cover profile
func (s *StdTester) WithCoverProfile(coverProfile string) *StdTester {
	s.coverProfile = coverProfile
	return s
}

// Test the project
func (s *StdTester) Test(c *Context) (err error) {
	var targets []string
	ctx := c.Cli.Context
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	gotest := buildkit.NewGoTest(targets...)
	gotest.WithCoverProfile(s.coverProfile)
	gotest.WithRace(true)

	cmd := gotest.Command(ctx)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Run testings")
	return cmd.Run()
}
