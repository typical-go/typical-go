package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestBuildTool_Validate(t *testing.T) {
	testcases := []struct {
		*typbuildtool.BuildTool
		expectedError string
	}{
		{
			BuildTool: typbuildtool.BuildSequences(typbuildtool.StandardBuild()),
		},
		{
			BuildTool:     typbuildtool.BuildSequences(),
			expectedError: "No build modules",
		},
		{
			BuildTool:     typbuildtool.BuildSequences(common.DummyValidator("some-error")),
			expectedError: "BuildTool: some-error",
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
