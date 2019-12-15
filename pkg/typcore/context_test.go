package typcore_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestContext_AllModule(t *testing.T) {
	ctx := typcore.Context{
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
	ctx := &typcore.Context{
		Name:    "some-name",
		Package: "some-package",
	}
	require.NoError(t, ctx.Validate())
	require.Equal(t, "0.0.1", ctx.Version)
	require.Equal(t, "*typcore.DefaultConfigLoader", reflect.TypeOf(ctx.ConfigLoader).String())
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		context typcore.Context
		errMsg  string
	}{
		{
			typcore.Context{Name: "some-name", Package: "some-package", AppModule: &dummyModule{Name: "App"}},
			"",
		},
		{
			typcore.Context{Package: "some-package"},
			"Context: Name can't be empty",
		},
		{
			typcore.Context{Name: "some-name"},
			"Context: Package can't be empty",
		},
		{
			typcore.Context{
				Name:    "some-name",
				Package: "some-package",
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linuxamd64"},
				},
			},
			"Context: Releaser: Target: Missing OS: Please make sure 'linuxamd64' using 'OS/ARCH' format",
		},
		{
			typcore.Context{
				Name:      "some-name",
				Package:   "some-package",
				AppModule: &dummyModule{Name: "App", err: errors.New("some-error")},
			},
			"Context: App: some-error",
		},
		{
			typcore.Context{
				Name:    "some-name",
				Package: "some-package",
				Modules: []interface{}{&dummyModule{Name: "Module", err: errors.New("some-error")}},
			},
			"Context: Module: some-error",
		},
	}
	for i, tt := range testcases {
		err := tt.context.Validate()
		if tt.errMsg == "" {
			require.NoError(t, err, i)
		} else {
			require.EqualError(t, err, tt.errMsg, i)
		}

	}
}

type dummyModule struct {
	Name string
	err  error
}

func (m dummyModule) Validate() error { return m.err }
