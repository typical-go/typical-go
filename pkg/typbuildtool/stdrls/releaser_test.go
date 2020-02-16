package stdrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool/stdrls"
)

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		*stdrls.Releaser
		errMsg string
	}{
		{
			stdrls.New().WithTarget(),
			"Missing 'Targets'",
		},
		{
			stdrls.New().WithTarget("invalid-target"),
			"Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Releaser.Validate(), tt.errMsg, i)
	}
}
