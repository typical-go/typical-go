package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
)

type (
	// Runner responsible to run
	Runner interface {
		Run(*Context) error
	}
	// Runners for composite run
	Runners []Runner
	// RunFn is runner function
	RunFn   func(*Context) error
	runImpl struct {
		fn RunFn
	}
	// StdRun standard run
	StdRun struct {
		Binary string
	}
)

//
// runImpl
//

// NewRunner return new instance of Run
func NewRunner(fn RunFn) Runner {
	return &runImpl{fn: fn}
}

func (r *runImpl) Run(c *Context) error {
	return r.fn(c)
}

//
// Runs
//

var _ Runner = (Runners)(nil)

// Run composite runner
func (r Runners) Run(c *Context) error {
	for _, runner := range r {
		if err := runner.Run(c); err != nil {
			return err
		}
	}
	return nil
}

//
// StdRun
//

var _ Runner = (*StdRun)(nil)

// Run for standard typical project
func (s *StdRun) Run(c *Context) error {
	if s.Binary == "" {
		s.Binary = fmt.Sprintf("bin/%s", c.Descriptor.Name)
	}

	return c.Execute(&execkit.Command{
		Name:   s.Binary,
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}
