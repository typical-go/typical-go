package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuildtool/typrls"
)

func TestTarget(t *testing.T) {
	testcases := []struct {
		typrls.Target
		os   string
		arch string
	}{
		{typrls.Target(""), "", ""},
		{typrls.Target("linux/amd"), "linux", "amd"},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.os, tt.OS(), i)
		require.Equal(t, tt.arch, tt.Arch(), i)
	}
}

func TestTarget_Validate(t *testing.T) {
	testcases := []struct {
		typrls.Target
		errMsg string
	}{
		{
			typrls.Target(""),
			"Can't be empty",
		},
		{
			typrls.Target("/amd"),
			"Missing OS: Please make sure '/amd' using 'OS/ARCH' format",
		},
		{
			typrls.Target("linux/"),
			"Missing Arch: Please make sure 'linux/' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.errMsg, i)
	}
}
