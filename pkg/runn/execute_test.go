package runn_test

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/runn"
)

func TestRun(t *testing.T) {
	testcases := []struct {
		stmts  []interface{}
		errMsg string
	}{
		{
			stmts: []interface{}{
				stdrun{"some-error"},
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
				stdrun{"some-error-1"},
				stdrun{"some-error-2"},
			},
			errMsg: "some-error-1",
		},
		{
			stmts: []interface{}{
				func() error { return errors.New("some-error-1") },
				func() error { return errors.New("some-error-2") },
			},
			errMsg: "some-error-1",
		},
		{
			stmts:  []interface{}{},
			errMsg: "",
		},
	}

	for i, tt := range testcases {
		err := runn.Run(tt.stmts...)
		if tt.errMsg == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}
	}
}

type stdrun struct {
	msg string
}

func (r stdrun) Run() error {
	if r.msg == "" {
		return nil
	}
	return errors.New(r.msg)
}
