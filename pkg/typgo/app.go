package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"go.uber.org/dig"
)

var (
	_ctors []*Constructor
	_dtors []*Destructor
)

type App struct {
	EntryPoint interface{}
}

// Provide constructor
func Provide(ctors ...*Constructor) {
	_ctors = append(_ctors, ctors...)
}

// Destroy destructor
func Destroy(dtors ...*Destructor) {
	_dtors = append(_dtors, dtors...)
}

// Run application
func (a *App) Run() error {
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		if _, err := LoadConfig(configFile); err != nil {
			return err
		}
	}

	di := dig.New()
	for _, c := range _ctors {
		if err := di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return err
		}
	}

	errs := common.GracefulRun(a.startFn(di), a.stopFn(di))
	return errs.Unwrap()
}

func (a *App) startFn(di *dig.Container) common.Fn {
	return func() (err error) {
		return di.Invoke(a.EntryPoint)
	}
}

func (a *App) stopFn(di *dig.Container) common.Fn {
	return func() (err error) {
		for _, dtor := range _dtors {
			if err = di.Invoke(dtor.Fn); err != nil {
				return
			}
		}
		return
	}
}
