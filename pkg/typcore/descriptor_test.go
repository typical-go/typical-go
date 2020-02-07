package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/stdrelease"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestContext_Validate_DefaultValue(t *testing.T) {
	desc := &typcore.Descriptor{
		Name:    "some-name",
		Package: "some-package",
		Build:   typbuild.New(),
	}
	require.NoError(t, desc.Validate())
	require.Equal(t, "0.0.1", desc.Version)
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		typcore.Descriptor
		errMsg string
	}{
		{
			typcore.Descriptor{Package: "some-package"},
			"Context: Name can't be empty",
		},
		{
			typcore.Descriptor{Name: "some-name"},
			"Context: Package can't be empty",
		},
		{
			typcore.Descriptor{
				Name:    "some-name",
				Package: "some-package",
				Build:   typbuild.New().WithRelease(stdrelease.New().WithTarget("linuxamd64")),
			},
			"Context: Build: Releaser: Target: Missing OS: Please make sure 'linuxamd64' using 'OS/ARCH' format",
		},
	}
	for i, tt := range testcases {
		err := tt.Validate()
		if tt.errMsg == "" {
			require.NoError(t, err, i)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}

	}
}
