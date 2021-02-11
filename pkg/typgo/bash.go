package typgo

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type (
	// Bash is wrapper to exec.Bash
	Bash struct {
		Name   string
		Args   []string
		Stdout io.Writer
		Stderr io.Writer
		Stdin  io.Reader
		Dir    string
		Env    []string
	}
	// Basher responsible to Bash
	Basher interface {
		Bash(extras ...string) *Bash
	}
	// RunExpectation is test expectation for typgo.RunBash
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

//
// Bash
//

var _ Basher = (*Bash)(nil)
var _ Action = (*Bash)(nil)
var _ fmt.Stringer = (*Bash)(nil)

// ExecCmd return exec.Cmd
func (b *Bash) ExecCmd(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, b.Name, b.Args...)
	cmd.Stdout = b.Stdout
	cmd.Stderr = b.Stderr
	cmd.Stdin = b.Stdin
	cmd.Dir = b.Dir
	cmd.Env = b.Env
	return cmd
}

// Bash return Bash
func (b *Bash) Bash(extras ...string) *Bash {
	return b
}

// Execute bash
func (b *Bash) Execute(c *Context) error {
	return c.Execute(b)
}

func (b Bash) String() string {
	var out strings.Builder
	fmt.Fprint(&out, b.Name)
	for _, arg := range b.Args {
		if strings.ContainsAny(arg, " ") {
			fmt.Fprintf(&out, " \"%s\"", arg)
		} else {
			fmt.Fprintf(&out, " %s", arg)
		}

	}
	return out.String()
}

//
// Run
//

var _mocker *runMocker

// RunBash the runner
func RunBash(ctx context.Context, cmder Basher) error {
	cmd := cmder.Bash()
	if _mocker != nil {
		return _mocker.run(ctx, cmd)
	}
	return cmd.ExecCmd(ctx).Run()
}

// PatchBash typgo.RunBash for testing purpose
func PatchBash(expectations []*RunExpectation) func(t *testing.T) {
	_mocker = &runMocker{
		expectations: expectations,
		ptr:          0,
		length:       len(expectations),
	}

	return func(t *testing.T) {
		if expectation := _mocker.expectation(); expectation != nil {
			t.Errorf("typgo-mock: missing call: %v", expectation)
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

func (r *runMocker) run(ctx context.Context, cmd *Bash) error {
	expc := r.expectation()
	if expc == nil {
		return fmt.Errorf("typgo-mock: no run expectation for \"%s\"", cmd.String())
	}

	if cmd.String() != expc.CommandLine {
		return fmt.Errorf("typgo-mock: \"%s\" should be \"%s\"", cmd.String(), expc.CommandLine)
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
