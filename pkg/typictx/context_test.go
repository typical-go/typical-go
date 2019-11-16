package typictx_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/typical-go/typical-go/pkg/typirelease"
)

func TestContext_AllModule(t *testing.T) {
	ctx := typictx.Context{
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
		context typictx.Context
		errMsg  string
	}{
		{
			typictx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Package:   "some-package",
				Releaser: typirelease.Releaser{
					Targets: []string{"linux/amd64"},
				},
			},
			"",
		},
		{
			typictx.Context{
				AppModule: dummyApp{},
				Package:   "some-package",
				Releaser: typirelease.Releaser{
					Targets: []string{"linux/amd64"},
				},
			},
			"Invalid Context: Name can't not empty",
		},
		{
			typictx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Releaser: typirelease.Releaser{
					Targets: []string{"linux/amd64"},
				},
			},
			"Invalid Context: Package can't not empty",
		},
		{
			typictx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Package:   "some-package",
			},
			"Release: Missing 'Targets'",
		},
		{
			typictx.Context{
				AppModule: dummyApp{},
				Name:      "some-name",
				Package:   "some-package",
				Releaser: typirelease.Releaser{
					Targets: []string{"invalid"},
				},
			},
			"Release: Missing '/' in target 'invalid'",
		},
	}
	for i, tt := range testcases {
		msg := fmt.Sprintf("Failed in case-%d", i)
		err := tt.context.Validate()
		if tt.errMsg == "" {
			require.NoError(t, err, msg)
		} else {
			require.EqualError(t, err, tt.errMsg, msg)
		}
	}
}

type dummyApp struct {
}

func (dummyApp) Run() (runFn interface{}) {
	return
}
