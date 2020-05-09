package common_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/common"
)

func TestValidate(t *testing.T) {
	testcases := []struct {
		testname string
		common.Validator
		expectedErr string
	}{
		{
			Validator:   nil,
			expectedErr: "nil",
		},
		{
			Validator:   &dummyValidator{"some-error"},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testname, func(t *testing.T) {
			err := common.Validate(tt.Validator)
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.expectedErr)
			}
		})
	}
}

type dummyValidator struct {
	errMsg string
}

// Validate return error
func (v *dummyValidator) Validate() error {
	if v.errMsg == "" {
		return nil
	}
	return errors.New(v.errMsg)
}
