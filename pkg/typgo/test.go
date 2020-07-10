package typgo

import (
	"fmt"
	"time"

	"github.com/typical-go/typical-go/pkg/execkit"
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
	// Testers for composite test
	Testers []Tester
	// StdTest is standard test
	StdTest struct {
		Timeout      time.Duration
		CoverProfile string
		Race         bool
	}
	// TestFn function
	TestFn     func(*Context) error
	testerImpl struct {
		fn TestFn
	}
)

var _ Tester = (*StdTest)(nil)

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

//
// Tests
//

var _ Tester = (Testers)(nil)

// Test composite
func (t Testers) Test(c *Context) (err error) {
	for _, test := range t {
		if err = test.Test(c); err != nil {
			return
		}
	}
	return
}
