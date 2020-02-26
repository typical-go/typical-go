package typtest

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// StdTester is standard tester
type StdTester struct{}

// New return new instance of StdTester
func New() *StdTester {
	return &StdTester{}
}

// Test the project
func (*StdTester) Test(ctx context.Context, c *Context) (err error) {

	log.Info("Run testings")

	var targets []string
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	args := []string{"test", "-coverprofile=cover.out", "-race"}
	args = append(args, targets...)

	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
