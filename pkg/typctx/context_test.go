package typctx_test

import (
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
		AppModule: dummyApp{},
	}
	require.Equal(t, 4, len(ctx.AllModule()))
	for i, module := range ctx.Modules {
		require.Equal(t, module, ctx.AllModule()[i])
	}
	require.Equal(t, ctx.AppModule, ctx.AllModule()[3])
}

func TestContext_Validate(t *testing.T) {
	testcases := []struct {
		context typctx.Context
		errMsg  string
	}{
		{
			typctx.Context{
				Name:      "some-name",
				Package:   "some-package",
				AppModule: dummyApp{},
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"",
		},
		{
			typctx.Context{
				Package:   "some-package",
				AppModule: dummyApp{},
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"Invalid Context: Name can't be empty",
		},
		{
			typctx.Context{
				Name:      "some-name",
				AppModule: dummyApp{},
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linux/amd64"},
				},
			},
			"Invalid Context: Package can't be empty",
		},
		{
			typctx.Context{
				Name:      "some-name",
				Package:   "some-package",
				AppModule: dummyApp{},
				Releaser: &typrls.Releaser{
					Targets: []typrls.Target{"linuxamd64"},
				},
			},
			"Releaser: Target: Missing OS: Please make sure 'linuxamd64' using 'OS/ARCH' format",
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
