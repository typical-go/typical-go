package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestValidate(t *testing.T) {
	testcases := []struct {
		*typbuildtool.StdBuild
		expected string
	}{
		{
			StdBuild: typbuildtool.StandardBuild().WithReleaseTargets("invalid-target"),
			expected: "Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.expected, i)
	}
}
