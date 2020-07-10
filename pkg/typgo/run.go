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

	// RunnerFn is runner function
	RunnerFn func(*Context) error

	runImpl struct {
		fn RunnerFn
	}

	// Runs for composite run
	Runs []Runner

	// StdRun standard run
	StdRun struct{}
)

var _ Runner = (*StdRun)(nil)

//
// runImpl
//

// NewRun return new instance of Run
func NewRun(fn RunnerFn) Runner {
	return &runImpl{fn: fn}
}

func (r *runImpl) Run(c *Context) error {
	return r.fn(c)
}

//
// Runs
//

// Run composite runner
func (r Runs) Run(c *Context) error {
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
