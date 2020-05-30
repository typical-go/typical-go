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
	// Tester responsible to test
	Tester interface {
		Test(*Context) error
	}

	// StdTest is standard test
	StdTest struct{}

	// TestFn function
	TestFn func(*Context) error

	// Tests for composite test
	Tests []Tester

	testerImpl struct {
		fn TestFn
	}
)

var _ Tester = (*StdTest)(nil)
var _ Tester = (Tests)(nil)

//
// testerImpl
//

// NewTest return new instance of test
func NewTest(fn TestFn) Tester {
	return &testerImpl{fn: fn}
}

func (t *testerImpl) Test(c *Context) error {
	return t.fn(c)
}

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

	return execute(c, gotest.Command())
}

//
// Tests
//

// Test composite
func (t Tests) Test(c *Context) (err error) {
	for _, test := range t {
		if err = test.Test(c); err != nil {
			return
		}
	}
	return
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
