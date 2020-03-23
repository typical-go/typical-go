package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestStdBuilder(t *testing.T) {
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.StandardBuild()
	})
	t.Run("SHOULD implement Cleaner", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.StandardBuild()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Tester = typbuildtool.StandardBuild()
	})
	t.Run("SHOULD implement Release", func(t *testing.T) {
		var _ typbuildtool.Releaser = typbuildtool.StandardBuild()
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typbuildtool.Commander = typbuildtool.StandardBuild()
	})
}

func TestValidate(t *testing.T) {
	testcases := []struct {
		*typbuildtool.StandardModule
		expected string
	}{
		{
			StandardModule: typbuildtool.StandardBuild().WithReleaseTargets(),
			expected:       "Missing 'Targets'",
		},
		{
			StandardModule: typbuildtool.StandardBuild().WithReleaseTargets("invalid-target"),
			expected:       "Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.expected, i)
	}
}
