package common_test

import (
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
			Validator:   common.DummyValidator("some-error"),
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
