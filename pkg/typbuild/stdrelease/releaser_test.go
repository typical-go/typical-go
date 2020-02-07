package stdrelease_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild/stdrelease"
)

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		*stdrelease.Releaser
		errMsg string
	}{
		{
			stdrelease.New().WithTarget(),
			"Missing 'Targets'",
		},
		{
			stdrelease.New().WithTarget("invalid-target"),
			"Target: Missing OS: Please make sure 'invalid-target' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Releaser.Validate(), tt.errMsg, i)
	}
}
