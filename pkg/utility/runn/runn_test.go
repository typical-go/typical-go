package runn_test

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/runn"
)

func TestExecute(t *testing.T) {
	testcases := []struct {
		stmts  []interface{}
		errMsg string
	}{
		{
			stmts: []interface{}{
				runner{"some-error"},
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
				runner{"some-error-1"},
				runner{"some-error-2"},
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
		err := runn.Execute(tt.stmts...)
		if tt.errMsg == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}
	}
}

type runner struct {
	msg string
}

func (r runner) Run() error {
	if r.msg == "" {
		return nil
	}
	return errors.New(r.msg)
}
