package typgo

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// Test responsible to test
	Test interface {
		Test(*Context) error
	}

	// StdTest is standard test
	StdTest struct{}
)

var _ Test = (*StdTest)(nil)

//
// StdTest
//

// Test standard
func (s *StdTest) Test(c *Context) (err error) {
	var targets []string

	for _, layout := range c.Descriptor.Layouts {
		targets = append(targets, fmt.Sprintf("./%s/...", layout))
	}

	if len(targets) < 1 {
		c.Info("Nothing to test")
		return
	}

	gotest := &buildkit.GoTest{
		Targets:      targets,
		Timeout:      typvar.TestTimeout,
		CoverProfile: typvar.TestCoverProfile,
		Race:         true,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
	}

	return execute(c, gotest)
}

//
// command
//

func cmdTest(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  c.ActionFn("TEST", test),
	}
}

func test(c *Context) error {
	if c.Test == nil {
		return errors.New("test is missing")
	}
	return c.Test.Test(c)
}
