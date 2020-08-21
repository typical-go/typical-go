package execkit

import (
	"context"
	"fmt"
	"testing"
)

var (
	_mocker *runMocker
)

type (
	// RunExpectation is test expectation for execkit.Run
	RunExpectation struct {
		CommandLine string
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
	cmd := cmder.Command()
	if _mocker != nil {
		return _mocker.run(ctx, cmd)
	}
	return cmd.ExecCmd(ctx).Run()
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

func (r *runMocker) run(ctx context.Context, cmd *Command) error {
	expc := r.expectation()
	if expc == nil {
		return fmt.Errorf("execkit-mock: no run expectation for \"%s\"", cmd.String())
	}

	if cmd.String() != expc.CommandLine {
		return fmt.Errorf("execkit-mock: \"%s\" should be \"%s\"", cmd.String(), expc.CommandLine)
	}

	if expc.OutputBytes != nil && cmd.Stdout != nil {
		if _, err := cmd.Stdout.Write(expc.OutputBytes); err != nil {
			return err
		}
	}
	if expc.ErrorBytes != nil && cmd.Stderr != nil {
		if _, err := cmd.Stderr.Write(expc.ErrorBytes); err != nil {
			return err
		}
	}

	return expc.ReturnError
}
