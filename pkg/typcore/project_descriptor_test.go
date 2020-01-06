package typcore_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestContext_AllModule(t *testing.T) {
	ctx := typcore.ProjectDescriptor{
		Modules: []interface{}{
			struct{}{},
			struct{}{},
			struct{}{},
		},
		AppModule: struct{}{},
	}
	require.Equal(t, 4, len(ctx.AllModule()))
	for i, module := range ctx.Modules {
		require.Equal(t, module, ctx.AllModule()[i])
	}
	require.Equal(t, ctx.AppModule, ctx.AllModule()[3])
}

func TestContext_Validate_DefaultValue(t *testing.T) {
	desc := &typcore.ProjectDescriptor{
		Name:    "some-name",
		Package: "some-package",
	}
	require.NoError(t, desc.Validate())
	require.Equal(t, "0.0.1", desc.Version)
	require.Equal(t, "*typcore.defaultConfigLoader", reflect.TypeOf(desc.ConfigLoader).String())
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
				Name:     "some-name",
				Package:  "some-package",
				Releaser: typrls.New().WithTarget("linuxamd64"),
			},
			"Context: Releaser: Target: Missing OS: Please make sure 'linuxamd64' using 'OS/ARCH' format",
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
