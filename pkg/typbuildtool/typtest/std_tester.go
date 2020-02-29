package typtest

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// StdTester is standard tester
type StdTester struct {
	coverProfile string
}

// New return new instance of StdTester
func New() *StdTester {
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
func (s *StdTester) Test(ctx context.Context, c *Context) (err error) {

	log.Info("Run testings")

	var targets []string
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	args := []string{"test", fmt.Sprintf("-coverprofile=%s", s.coverProfile), "-race"}
	args = append(args, targets...)

	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
