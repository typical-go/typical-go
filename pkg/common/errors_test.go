package common_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/common"
)

func TestErrors(t *testing.T) {
	testcases := []struct {
		testName string
		*common.Errors
		sep   string
		msg   string
		error error
	}{
		{
			Errors: new(common.Errors).Append(
				errors.New("error1"),
				errors.New("error2"),
				errors.New("error3"),
			),
			sep:   "+",
			msg:   "error1+error2+error3",
			error: errors.New("error1; error2; error3"),
		},
		{
			Errors: new(common.Errors).
				Append(errors.New("error1")).
				Append(errors.New("error2")),
			sep:   "|",
			msg:   "error1|error2",
			error: errors.New("error1; error2"),
		},
		{
			Errors: &common.Errors{},
			msg:    "",
			error:  nil,
		},
	}
	for i, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.msg, tt.Join(tt.sep), i)
			if err := tt.Unwrap(); err != nil {
				require.EqualError(t, err, tt.error.Error(), i)
			} else {
				require.NoError(t, err, i)
			}
		})
	}
}
