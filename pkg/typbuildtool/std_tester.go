package typbuildtool

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/exor"
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
func (s *StdTester) Test(c *BuildContext) (err error) {
	var targets []string
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	gotest := exor.NewGoTest(targets...).
		WithCoverProfile(s.coverProfile).
		WithRace(true).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr)

	log.Info("Run testings")
	return gotest.Execute(c.Cli.Context)
}
