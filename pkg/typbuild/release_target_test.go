package typbuild_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

func TestTarget(t *testing.T) {
	testcases := []struct {
		typbuild.ReleaseTarget
		os   string
		arch string
	}{
		{typbuild.ReleaseTarget(""), "", ""},
		{typbuild.ReleaseTarget("linux/amd"), "linux", "amd"},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.os, tt.OS(), i)
		require.Equal(t, tt.arch, tt.Arch(), i)
	}
}

func TestTarget_Validate(t *testing.T) {
	testcases := []struct {
		typbuild.ReleaseTarget
		errMsg string
	}{
		{
			typbuild.ReleaseTarget(""),
			"Can't be empty",
		},
		{
			typbuild.ReleaseTarget("/amd"),
			"Missing OS: Please make sure '/amd' using 'OS/ARCH' format",
		},
		{
			typbuild.ReleaseTarget("linux/"),
			"Missing Arch: Please make sure 'linux/' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		require.EqualError(t, tt.Validate(), tt.errMsg, i)
	}
}
