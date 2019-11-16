package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestReleaseTarget(t *testing.T) {
	testcases := []struct {
		typrls.ReleaseTarget
		os   string
		arch string
	}{
		{typrls.ReleaseTarget(""), "", ""},
		{typrls.ReleaseTarget("linux/amd"), "linux", "amd"},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.os, tt.OS(), i)
		require.Equal(t, tt.arch, tt.Arch(), i)
	}
}

func TestReleaseTarget_Validate(t *testing.T) {
	testcases := []struct {
		typrls.ReleaseTarget
		errMsg string
	}{
		{
			typrls.ReleaseTarget(""),
			"Can't be empty",
		},
		{
			typrls.ReleaseTarget("/amd"),
			"Missing OS: Please make sure '/amd' using 'OS/ARCH' format",
		},
		{
			typrls.ReleaseTarget("linux/"),
			"Missing Arch: Please make sure 'linux/' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.errMsg, i)
	}
}
