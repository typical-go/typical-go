package typbuild_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

func TestBuildTool_Validate(t *testing.T) {
	testcases := []struct {
		*typbuild.BuildTool
		expectedError string
	}{
		{
			BuildTool: &typbuild.BuildTool{
				BuildSequences: []interface{}{
					typbuild.StandardBuild(),
				},
			},
		},
		{
			BuildTool: &typbuild.BuildTool{
				BuildSequences: []interface{}{
					typbuild.StandardBuild(),
				},
			},
			expectedError: "No build-sequence",
		},
		{
			BuildTool: &typbuild.BuildTool{
				BuildSequences: []interface{}{
					typbuild.StandardBuild(),
				},
			},
			expectedError: "build-seq-error",
		},
		{
			BuildTool: &typbuild.BuildTool{
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
	typbuild.SimpleUtility
	errMsg string
}
