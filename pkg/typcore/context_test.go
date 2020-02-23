package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestTypicalContext_Validate(t *testing.T) {
	testcases := []struct {
		*typcore.TypicalContext
		expectedError string
	}{
		{
			TypicalContext: &typcore.TypicalContext{},
			expectedError:  "TypicalContext: ModulePackage can't be empty",
		},
		{
			TypicalContext: &typcore.TypicalContext{ModulePackage: "some-package"},
		},
	}

	for _, tt := range testcases {
		err := common.Validate(tt.TypicalContext)
		if tt.expectedError == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.expectedError)
		}
	}

}
