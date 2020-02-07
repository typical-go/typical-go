package stdrelease_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/stdrelease"
)

func TestTarget(t *testing.T) {
	testcases := []struct {
		stdrelease.Target
		os   string
		arch string
	}{
		{stdrelease.Target(""), "", ""},
		{stdrelease.Target("linux/amd"), "linux", "amd"},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.os, tt.OS(), i)
		require.Equal(t, tt.arch, tt.Arch(), i)
	}
}

func TestTarget_Validate(t *testing.T) {
	testcases := []struct {
		stdrelease.Target
		errMsg string
	}{
		{
			stdrelease.Target(""),
			"Can't be empty",
		},
		{
			stdrelease.Target("/amd"),
			"Missing OS: Please make sure '/amd' using 'OS/ARCH' format",
		},
		{
			stdrelease.Target("linux/"),
			"Missing Arch: Please make sure 'linux/' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.errMsg, i)
	}
}
