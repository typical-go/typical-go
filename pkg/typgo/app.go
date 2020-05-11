package typgo

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"go.uber.org/dig"
)

func launchApp(d *Descriptor) (err error) {
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		_, err = typcfg.Load(configFile)
	}

	di := dig.New()
	if err = provide(di, d); err != nil {
		return
	}

	startGracefuly(start(di, d), stop(di))
	return
}

func provide(di *dig.Container, d *Descriptor) (err error) {
	if err = di.Provide(func() *Descriptor { return d }); err != nil {
		return
	}
	for _, c := range _ctors {
		if err = di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return
		}
	}
	return
}

func start(di *dig.Container, d *Descriptor) func() error {
	return func() (err error) {
		return di.Invoke(d.EntryPoint)
	}
}

func stop(di *dig.Container) func() error {
	return func() (err error) {
		for _, dtor := range _dtors {
			if err = di.Invoke(dtor.Fn); err != nil {
				return
			}
		}
		return
	}
}

func startGracefuly(startFn func() error, stopFn func() error) (errs common.Errors) {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		if err := startFn(); err != nil {
			// NOTE: if startFn got error, it should still execute stopFn
			errs.Append(err)
		}
	}()
	<-gracefulStop
	if stopFn != nil {
		if err := stopFn(); err != nil {
			errs.Append(err)
		}
	}
	return
}
