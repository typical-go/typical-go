package typgo

import (
	"errors"
	"fmt"
	"time"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/urfave/cli/v2"
)

const (
	defaultTestTimeout      = 25 * time.Second
	defaultTestCoverProfile = "cover.out"
)

type (
	// Tester responsible to test
	Tester interface {
		Test(*Context) error
	}

	// StdTest is standard test
	StdTest struct {
		Timeout      time.Duration
		CoverProfile string
		Race         bool
	}

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

	if len(c.Descriptor.Layouts) < 1 {
		c.Info("Nothing to test")
		return
	}

	return c.Execute(&buildkit.GoTest{
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
