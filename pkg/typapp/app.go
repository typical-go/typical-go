package typapp

import (
	"os"
	"os/signal"
	"syscall"

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

	exitSig := make(chan os.Signal)
	signal.Notify(exitSig, syscall.SIGTERM)
	signal.Notify(exitSig, syscall.SIGINT)

	var errs common.Errors

	go func() {
		defer func() { exitSig <- syscall.SIGTERM }()
		if err := di.Invoke(a.EntryPoint); err != nil {
			errs.Append(err)
		}
	}()
	<-exitSig

	for _, dtor := range a.Dtors {
		if err := di.Invoke(dtor.Fn); err != nil {
			errs.Append(err)
		}
	}
	return errs.Unwrap()
}
