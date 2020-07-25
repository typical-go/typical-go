package execkit

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

var (
	_mocker *runMocker
)

type (
	// RunExpectation is test expectation for execkit.Run
	RunExpectation struct {
		CommandLine []string
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
	cmd := cmder.Command()
	if expc == nil {
		return fmt.Errorf("execkit-mock: no run expectation for [%s %s]",
			cmd.Name, strings.Join(cmd.Args, " "))
	}

	if !expc.match(cmd) {
		return fmt.Errorf("execkit-mock: [%s %s] should be %v",
			cmd.Name, strings.Join(cmd.Args, " "), expc.CommandLine)
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

func (r *RunExpectation) match(c *Command) bool {
	if len(c.Args)+1 != len(r.CommandLine) {
		return false
	}

	if r.CommandLine[0] != c.Name {
		return false
	}

	for i, arg := range c.Args {
		if arg != r.CommandLine[i+1] {
			return false
		}
	}

	return true
}
