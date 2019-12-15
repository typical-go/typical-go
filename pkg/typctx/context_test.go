package typctx_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestContext_AllModule(t *testing.T) {
	ctx := typctx.Context{
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
	ctx := &typctx.Context{
		Name:    "some-name",
		Package: "some-package",
	}
	require.NoError(t, ctx.Validate())
	require.Equal(t, "0.0.1", ctx.Version)
	require.Equal(t, "*typobj.defaultLoader", reflect.TypeOf(ctx.ConfigLoader).String())
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		context typctx.Context
		errMsg  string
	}{
		{
			typctx.Context{Name: "some-name", Package: "some-package", AppModule: &dummyModule{Name: "App"}},
			"",
		},
		{
			typctx.Context{Package: "some-package"},
			"Context: Name can't be empty",
		},
		{
			typctx.Context{Name: "some-name"},
			"Context: Package can't be empty",
		},
		{
			typctx.Context{
				Name:    "some-name",
				Package: "some-package",
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linuxamd64"},
				},
			},
			"Context: Releaser: Target: Missing OS: Please make sure 'linuxamd64' using 'OS/ARCH' format",
		},
		{
			typctx.Context{
				Name:      "some-name",
				Package:   "some-package",
				AppModule: &dummyModule{Name: "App", err: errors.New("some-error")},
			},
			"Context: App: some-error",
		},
		{
			typctx.Context{
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
