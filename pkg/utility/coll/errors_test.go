package coll_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestErrors(t *testing.T) {
	testcases := []struct {
		coll.Errors
		errors []error
		msg    string
		error  error
	}{
		{
			Errors: new(coll.Errors).Append(
				errors.New("error1"),
				errors.New("error2"),
				errors.New("error3"),
			),
			errors: []error{
				errors.New("error1"),
				errors.New("error2"),
				errors.New("error3"),
			},
			msg:   "error1; error2; error3",
			error: errors.New("error1; error2; error3"),
		},
		{
			Errors: new(coll.Errors).
				Append(errors.New("error1")).
				Append(errors.New("error2")),
			errors: []error{
				errors.New("error1"),
				errors.New("error2"),
			},
			msg:   "error1; error2",
			error: errors.New("error1; error2"),
		},
		{
			Errors: coll.Errors{},
			errors: []error{},
			msg:    "",
			error:  nil,
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.errors, tt.Errors)
		require.Equal(t, tt.msg, tt.Error(), i)
		if err := tt.ToError(); err != nil {
			require.EqualError(t, err, tt.error.Error(), i)
		} else {
			require.NoError(t, err, i)
		}
	}
}
