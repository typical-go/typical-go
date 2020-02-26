package common_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/common"
)

func TestValidate(t *testing.T) {
	testcases := []struct {
		common.Validator
		expectedErr string
	}{
		{},
		{
			Validator:   &validator{errors.New("some-error")},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		err := common.Validate(tt.Validator)
		if tt.expectedErr == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.expectedErr)
		}
	}

}

type validator struct {
	err error
}

func (v *validator) Validate() error {
	return v.err
}
