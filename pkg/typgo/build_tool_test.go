package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestBuildTool_Validate(t *testing.T) {
	testcases := []struct {
		*typgo.BuildTool
		expectedError string
	}{
		{
			BuildTool: &typgo.BuildTool{
				BuildSequences: []interface{}{
					typgo.StandardBuild(),
				},
			},
		},
		{
			BuildTool: &typgo.BuildTool{
				BuildSequences: []interface{}{
					typgo.StandardBuild(),
				},
			},
			expectedError: "No build-sequence",
		},
		{
			BuildTool: &typgo.BuildTool{
				BuildSequences: []interface{}{
					typgo.StandardBuild(),
				},
			},
			expectedError: "build-seq-error",
		},
		{
			BuildTool: &typgo.BuildTool{
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
	typgo.SimpleUtility
	errMsg string
}
