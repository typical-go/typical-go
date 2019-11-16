package typctx_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestContext_AllModule(t *testing.T) {
	ctx := typctx.Context{
		AppModule: dummyApp{},
		Modules: []interface{}{
			struct{}{},
			struct{}{},
			struct{}{},
		},
	}
	require.Equal(t, 4, len(ctx.AllModule()))
	require.Equal(t, ctx.AppModule, ctx.AllModule()[0])
	for i, module := range ctx.Modules {
		require.Equal(t, module, ctx.AllModule()[i+1])
	}
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		context typctx.Context
		errMsg  string
	}{
		{
			typctx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Package:   "some-package",
				Releaser: typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"",
		},
		{
			typctx.Context{
				AppModule: dummyApp{},
				Package:   "some-package",
				Releaser: typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"Invalid Context: Name can't not empty",
		},
		{
			typctx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Releaser: typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"Invalid Context: Package can't not empty",
		},
		{
			typctx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Package:   "some-package",
			},
			"Releaser: Missing 'Targets'",
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

type dummyApp struct {
}

func (dummyApp) Run() (runFn interface{}) {
	return
}
