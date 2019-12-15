package typrls_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestReleaser(t *testing.T) {
	releaser := &typrls.Releaser{}
	require.True(t, typcore.IsValidator(releaser))
}

func TestReleaser_Validate(t *testing.T) {
	testcases := []struct {
		typrls.Releaser
		errMsg string
	}{
		{
			typrls.Releaser{},
			"Missing 'Targets'",
		},
		{
			typrls.Releaser{
				Targets: []typrls.Target{"invalid"},
			},
			"Target: Missing OS: Please make sure 'invalid' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Releaser.Validate(), tt.errMsg, i)
	}
}
