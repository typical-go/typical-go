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
	// RunExpectation is test expectation for execkit.Run
	RunExpectation struct {
		Ctx         context.Context
		Command     Commander
		ReturnError error
	}
	runMocker struct {
		expectations []*RunExpectation
		ptr          int
		length       int
	}
)

// Run the runner
func Run(ctx context.Context, cmder Commander) error {
	if _mocker != nil {
		return _mocker.run(ctx, cmder)
	}
	return cmder.Command().Run(ctx)
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

func (r *runMocker) run(ctx context.Context, cmder Commander) error {
	expectation := r.expectation()
	if expectation == nil {
		return errors.New("execkit-mock: no run expectation")
	}
	if expectation.Ctx != nil && !reflect.DeepEqual(expectation.Ctx, ctx) {
		return errors.New("execkit-mock: context not match")
	}
	if expectation.Command != nil {
		cmd1 := cmder.Command()
		cmd2 := expectation.Command.Command()
		if cmd1.Name != cmd2.Name || !reflect.DeepEqual(cmd1.Args, cmd2.Args) {
			return fmt.Errorf("execkit-mock: command not match: {%s %v} != {%s %v}",
				cmd1.Name, cmd1.Args, cmd2.Name, cmd2.Args)
		}

	}
	return expectation.ReturnError
}
