package typirelease_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typirelease"
)

func TestReleaseTarget(t *testing.T) {
	testcases := []struct {
		typirelease.ReleaseTarget
		os   string
		arch string
	}{
		{typirelease.ReleaseTarget(""), "", ""},
		{typirelease.ReleaseTarget("linux/amd"), "linux", "amd"},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.os, tt.OS(), i)
		require.Equal(t, tt.arch, tt.Arch(), i)
	}
}

func TestReleaseTarget_Validate(t *testing.T) {
	testcases := []struct {
		typirelease.ReleaseTarget
		errMsg string
	}{
		{
			typirelease.ReleaseTarget(""),
			"Can't be empty",
		},
		{
			typirelease.ReleaseTarget("/amd"),
			"Missing OS: Please make sure '/amd' using 'OS/ARCH' format",
		},
		{
			typirelease.ReleaseTarget("linux/"),
			"Missing Arch: Please make sure 'linux/' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.errMsg, i)
	}
}
