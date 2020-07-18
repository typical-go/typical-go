package execkit

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

var (
	_mocker *runMocker
)

type (
	// Runner responsble to run
	Runner interface {
		Run(context.Context) error
	}
	// RunFn run function
	RunFn      func(context.Context) error
	runnerImpl struct {
		fn RunFn
	}
	// RunExpectation is test expectation for execkit.Run
	RunExpectation struct {
		Ctx         context.Context
		Runner      Runner
		ReturnError error
	}
	runMocker struct {
		expectations []*RunExpectation
		ptr          int
		length       int
	}
)

// Run the runner
func Run(ctx context.Context, runner Runner) error {
	if _mocker != nil {
		return _mocker.run(ctx, runner)
	}
	return runner.Run(ctx)
}

//
// runnerImpl
//

// NewRunner return new instance of runners
func NewRunner(fn RunFn) Runner {
	return &runnerImpl{fn: fn}
}

func (r *runnerImpl) Run(ctx context.Context) error {
	return r.fn(ctx)
}

//
// Patch
//

// Patch execkit.Run for testing purpose
func Patch(expectations []*RunExpectation) func(t *testing.T) {
	_mocker = &runMocker{
		expectations: expectations,
		ptr:          0,
		length:       len(expectations),
	}

	return func(t *testing.T) {
		if expectation := _mocker.expectation(); expectation != nil {
			t.Errorf("execkit-mock: missing call: %v", expectation)
			t.FailNow()
		}
		_mocker = nil
	}
}

func (r *runMocker) expectation() *RunExpectation {
	if r.ptr < r.length {
		expect := r.expectations[r.ptr]
		r.ptr++
		return expect
	}
	return nil
}

func (r *runMocker) run(ctx context.Context, runner Runner) error {
	expectation := r.expectation()
	if expectation == nil {
		return errors.New("execkit-mock: no run expectation")
	}
	if expectation.Ctx != nil && !reflect.DeepEqual(expectation.Ctx, ctx) {
		return errors.New("execkit-mock: context not match")
	}
	if expectation.Runner != nil && !reflect.DeepEqual(expectation.Runner, runner) {
		return fmt.Errorf("execkit-mock: runner not match: %v", runner)
	}
	return expectation.ReturnError
}
