package runn_test

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/runn"
)

func TestExecutor_All(t *testing.T) {
	testcases := []struct {
		stmts           []interface{}
		stopableErrMsg  string
		unstopaleErrMsg string
	}{
		{
			stmts: []interface{}{
				runner{"some-error"},
				exec.Command("wrong-command", "bad-argument"),
			},
			stopableErrMsg:  "some-error",
			unstopaleErrMsg: "some-error; exec: \"wrong-command\": executable file not found in $PATH",
		},
		{
			stmts: []interface{}{
				errors.New("error1"),
				errors.New("error2"),
			},
			stopableErrMsg:  "Statement-0: Invalid: *errors.errorString",
			unstopaleErrMsg: "Statement-0: Invalid: *errors.errorString",
		},
		{
			stmts: []interface{}{
				runner{"some-error-1"},
				runner{"some-error-2"},
			},
			stopableErrMsg:  "some-error-1",
			unstopaleErrMsg: "some-error-1; some-error-2",
		},
	}

	for i, tt := range testcases {
		err := runn.Executor{StopWhenError: true}.Execute(tt.stmts...)
		if tt.stopableErrMsg == "" {
			require.NoError(t, err, fmt.Sprintf("stopable-%d", i))
		} else {
			require.EqualError(t, err, tt.stopableErrMsg, fmt.Sprintf("stopable-%d", i))
		}
		err = runn.Executor{StopWhenError: false}.Execute(tt.stmts...)
		if tt.unstopaleErrMsg == "" {
			require.NoError(t, err, fmt.Sprintf("unstopable-%d", i))
		} else {
			require.EqualError(t, err, tt.unstopaleErrMsg, i, fmt.Sprintf("unstopable-%d", i))
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
