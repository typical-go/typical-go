package typirelease_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typirelease"
)

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		typirelease.Releaser
		errMsg string
	}{
		{
			typirelease.Releaser{},
			"Missing 'Targets'",
		},
		{
			typirelease.Releaser{
				Targets: []typirelease.ReleaseTarget{"invalid"},
			},
			"Target: Missing OS: Please make sure 'invalid' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Releaser.Validate(), tt.errMsg, i)
	}
}
