package runn_test

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/runn"
)

type RunnerImplementationWithError struct{}

func (i RunnerImplementationWithError) Run() error { return errors.New("some-runner-error") }

type RunnerImplementationNoError struct{}

func (i RunnerImplementationNoError) Run() error { return nil }

func TestExecutor_All(t *testing.T) {
	testcases := []struct {
		stopWhenError bool
		stmts         []interface{}
		err           error
	}{
		{
			false,
			[]interface{}{
				errors.New("error1"),
				RunnerImplementationWithError{},
				errors.New("error2"),
				exec.Command("wrong-command", "bad-argument"),
			},
			errors.New("error1; some-runner-error; error2; exec: \"wrong-command\": executable file not found in $PATH"),
		},
		{
			true,
			[]interface{}{
				errors.New("error1"),
				errors.New("unreachable-error"),
			},
			errors.New("error1"),
		},
		{
			true,
			[]interface{}{
				RunnerImplementationWithError{},
				errors.New("unreachable-error"),
			},
			errors.New("some-runner-error"),
		},
		{
			true,
			[]interface{}{
				exec.Command("wrong-command", "bad-argument"),
				errors.New("unreachable-error"),
			},
			errors.New("exec: \"wrong-command\": executable file not found in $PATH"),
		},
		{
			false,
			[]interface{}{},
			nil,
		},

		{
			false,
			[]interface{}{
				RunnerImplementationNoError{},
				nil,
				exec.Command("echo", "hello", "world"),
			},
			nil,
		},
	}

	for _, tt := range testcases {
		err := runn.Executor{
			StopWhenError: tt.stopWhenError,
		}.Execute(tt.stmts...)

		if tt.err == nil {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.err.Error())
		}
	}
}
