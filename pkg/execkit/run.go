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
		Command     Commander
		OutputBytes []byte
		ErrorBytes  []byte
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
	expc := r.expectation()
	if expc == nil {
		return errors.New("execkit-mock: no run expectation")
	}
	if expc.Command != nil {
		if err := check(cmder, expc.Command); err != nil {
			return err
		}

		cmd := cmder.Command()
		if cmd.Stdout != nil {
			if _, err := cmd.Stdout.Write(expc.OutputBytes); err != nil {
				return err
			}
		}
		if cmd.Stderr != nil {
			if _, err := cmd.Stderr.Write(expc.ErrorBytes); err != nil {
				return err
			}
		}

	}
	return expc.ReturnError
}

func check(c1, c2 Commander) error {
	cmd1 := c1.Command()
	cmd2 := c2.Command()
	if cmd1.Name != cmd2.Name || !reflect.DeepEqual(cmd1.Args, cmd2.Args) {
		return fmt.Errorf("execkit-mock: command not match: {%s %v} != {%s %v}",
			cmd1.Name, cmd1.Args, cmd2.Name, cmd2.Args)
	}
	return nil
}
