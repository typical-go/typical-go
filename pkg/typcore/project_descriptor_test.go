package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestContext_Validate_DefaultValue(t *testing.T) {
	desc := &typcore.ProjectDescriptor{
		Name:    "some-name",
		Package: "some-package",
	}
	require.NoError(t, desc.Validate())
	require.Equal(t, "0.0.1", desc.Version)
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		typcore.ProjectDescriptor
		errMsg string
	}{
		{
			typcore.ProjectDescriptor{Package: "some-package"},
			"Context: Name can't be empty",
		},
		{
			typcore.ProjectDescriptor{Name: "some-name"},
			"Context: Package can't be empty",
		},
		{
			typcore.ProjectDescriptor{
				Name:    "some-name",
				Package: "some-package",
				Build:   typcore.NewBuild().WithRelease(typrls.New().WithTarget("linuxamd64")),
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
