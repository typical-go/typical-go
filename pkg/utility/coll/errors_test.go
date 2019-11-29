package coll_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestErrors(t *testing.T) {
	testcases := []struct {
		*coll.Errors
		slice []error
		sep   string
		msg   string
		error error
	}{
		{
			Errors: new(coll.Errors).Append(
				errors.New("error1"),
				errors.New("error2"),
				errors.New("error3"),
			),
			slice: []error{
				errors.New("error1"),
				errors.New("error2"),
				errors.New("error3"),
			},
			sep:   "+",
			msg:   "error1+error2+error3",
			error: errors.New("error1; error2; error3"),
		},
		{
			Errors: new(coll.Errors).
				Append(errors.New("error1")).
				Append(errors.New("error2")),
			slice: []error{
				errors.New("error1"),
				errors.New("error2"),
			},
			sep:   "|",
			msg:   "error1|error2",
			error: errors.New("error1; error2"),
		},
		{
			Errors: &coll.Errors{},
			slice:  []error{},
			msg:    "",
			error:  nil,
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.Errors.Slice())
		require.Equal(t, tt.msg, tt.Join(tt.sep), i)
		if err := tt.Unwrap(); err != nil {
			require.EqualError(t, err, tt.error.Error(), i)
		} else {
			require.NoError(t, err, i)
		}
	}
}
