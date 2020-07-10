package typgo

import (
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
	StdRun struct{}
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
func (*StdRun) Run(c *Context) error {
	return c.Execute(&execkit.Command{
		Name:   AppBin(c.Descriptor.Name),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}
