package runnerkit_test

import (
	"context"
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/runnerkit"
)

func TestRun(t *testing.T) {
	testcases := []struct {
		stmts  []interface{}
		errMsg string
	}{
		{
			stmts: []interface{}{
				errorRunner{"some-error"},
				exec.Command("wrong-command", "bad-argument"),
			},
			errMsg: "some-error",
		},
		{
			stmts: []interface{}{
				errors.New("error1"),
				errors.New("error2"),
			},
			errMsg: "Statement-0: Invalid: *errors.errorString",
		},
		{
			stmts: []interface{}{
				errorRunner{"some-error-1"},
				errorRunner{"some-error-2"},
			},
			errMsg: "some-error-1",
		},
		{
			stmts: []interface{}{
				func(context.Context) error { return errors.New("some-error-1") },
				func(context.Context) error { return errors.New("some-error-2") },
			},
			errMsg: "some-error-1",
		},
		{
			stmts:  []interface{}{},
			errMsg: "",
		},
	}

	for i, tt := range testcases {
		err := runnerkit.Run(context.Background(), tt.stmts...)
		if tt.errMsg == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}
	}
}

type errorRunner struct {
	msg string
}

func (r errorRunner) Run(ctx context.Context) error {
	if r.msg == "" {
		return nil
	}
	return errors.New(r.msg)
}
