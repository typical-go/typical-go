package typcli

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typmodule"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// AppCli is cli for context
type AppCli struct {
	*typctx.Context
}

// Action to return action function
func (c *AppCli) Action(fn interface{}) func(ctx *cli.Context) error {
	return preparedAction(fn, c.Context)
}

func preparedAction(fn interface{}, c *typctx.Context) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		gracefulStop := make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		go func() {
			<-gracefulStop
			if err = destroyAll(di, c); err != nil {
				log.Fatal(err.Error())
			}
			os.Exit(0) // NOTE: Make sure the application is exit
		}()
		if err = provideAll(di, c); err != nil {
			return
		}
		if err = prepareAll(di, c); err != nil {
			return
		}
		return di.Invoke(fn)
	}
}

func provideAll(di *dig.Container, ctx *typctx.Context) (err error) {
	if err = provideLoader(di, ctx); err != nil {
		return
	}
	if err = provide(di, ctx.Constructors...); err != nil {
		return
	}
	for _, module := range ctx.AllModule() {
		if err = provideConfigFn(di, module); err != nil {
			return
		}
		if provider, ok := module.(typmodule.Provider); ok {
			if err = provide(di, provider.Provide()...); err != nil {
				return
			}
		}
	}
	return
}

func prepareAll(di *dig.Container, ctx *typctx.Context) (err error) {
	for _, module := range ctx.AllModule() {
		if preparer, ok := module.(typmodule.Preparer); ok {
			if err = invoke(di, preparer.Prepare()...); err != nil {
				return
			}
		}
	}
	return
}

func destroyAll(di *dig.Container, ctx *typctx.Context) (err error) {
	for _, module := range ctx.AllModule() {
		if destroyer, ok := module.(typmodule.Destroyer); ok {
			if err = invoke(di, destroyer.Destroy()...); err != nil {
				return
			}
		}
	}
	return
}

func invoke(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Invoke(fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Provide(fn); err != nil {
			return
		}
	}
	return
}

func loaderFn(c *typctx.Context) interface{} {
	return func() typcfg.Loader {
		return c.ConfigLoader
	}
}

func provideLoader(di *dig.Container, ctx *typctx.Context) (err error) {
	if ctx.ConfigLoader != nil {
		if err = provide(di, loaderFn(ctx)); err != nil {
			return
		}
	}
	return
}

func provideConfigFn(di *dig.Container, v interface{}) (err error) {
	if configurer, ok := v.(typcfg.Configurer); ok {
		_, _, loadFn := configurer.Configure()
		if err = provide(di, loadFn); err != nil {
			return
		}
	}
	return
}
