package typapp

import (
	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Invoke function with Dependency Injection
func (a *App) Invoke(actx *typcore.AppContext, c *cli.Context, fn interface{}) (err error) {
	di := dig.New()
	if c != nil {
		if err = di.Provide(func() *cli.Context { return c }); err != nil {
			return
		}
	}
	if actx.Configuration != nil {
		// provide configuration to dependency-injection container
		if err = provide(di, actx.Configuration.Provide()...); err != nil {
			return
		}
	}
	// provide registered function in descriptor to dependency-injection container
	if err = provide(di, a.Provide()...); err != nil {
		return
	}
	// invoke preparation as register in descriptor
	if err = invoke(di, a.Prepare()...); err != nil {
		return
	}

	startFn := func() error {
		return di.Invoke(fn)
	}
	stopFn := func() error {
		return invoke(di, a.Destroy()...)
	}
	for _, err := range common.StartGracefully(startFn, stopFn) {
		log.Error(err.Error())
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
