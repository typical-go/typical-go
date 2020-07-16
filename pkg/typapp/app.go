package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"go.uber.org/dig"
)

type (
	// App of typical program
	App struct {
		EntryPoint interface{}
		Ctors      []*Constructor
		Dtors      []*Destructor
	}
)

// Run application
func (a *App) Run() error {
	di := dig.New()
	for _, c := range a.Ctors {
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
		for _, dtor := range a.Dtors {
			if err = di.Invoke(dtor.Fn); err != nil {
				return
			}
		}
		return
	}
}
