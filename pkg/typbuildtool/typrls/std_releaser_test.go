package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool/typrls"
)

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		*typrls.StdReleaser
		expected string
	}{
		{
			StdReleaser: typrls.New().WithTarget(),
			expected:    "Missing 'Targets'",
		},
		{
			StdReleaser: typrls.New().WithTarget("invalid-target"),
			expected:    "Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.expected, i)
	}
}
