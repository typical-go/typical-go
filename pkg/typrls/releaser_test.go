package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		*typrls.Releaser
		errMsg string
	}{
		{
			typrls.New().WithTarget(),
			"Missing 'Targets'",
		},
		{
			typrls.New().WithTarget("invalid-target"),
			"Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Releaser.Validate(), tt.errMsg, i)
	}
}
