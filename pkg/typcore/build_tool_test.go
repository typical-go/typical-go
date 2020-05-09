package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestBuildTool_Validate(t *testing.T) {
	testcases := []struct {
		*typcore.BuildTool
		expectedError string
	}{
		{
			BuildTool: &typcore.BuildTool{
				BuildSequences: []interface{}{
					typcore.StandardBuild(),
				},
			},
		},
		{
			BuildTool: &typcore.BuildTool{
				BuildSequences: []interface{}{
					typcore.StandardBuild(),
				},
			},
			expectedError: "No build-sequence",
		},
		{
			BuildTool: &typcore.BuildTool{
				BuildSequences: []interface{}{
					typcore.StandardBuild(),
				},
			},
			expectedError: "build-seq-error",
		},
		{
			BuildTool: &typcore.BuildTool{
				BuildSequences: []interface{}{
					struct{}{},
				},
				Utility: &utilityWithErrors{errMsg: "utility-error"},
			},
			expectedError: "utility-error",
		},
	}

	for _, tt := range testcases {
		if err := tt.Validate(); err != nil {
			require.EqualError(t, err, tt.expectedError)
		} else {
			require.NoError(t, err)
		}
	}
}

type utilityWithErrors struct {
	typcore.SimpleUtility
	errMsg string
}
