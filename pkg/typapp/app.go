package typapp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

// Run the entry points
func Run(entryPoint interface{}) error {
	app := &App{
		EntryPoint: entryPoint,
		Ctors:      GetCtors(),
		Dtors:      GetDtors(),
	}
	return app.Run()
}

//
// app
//

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

	var err error

	go func() {
		defer func() { exitSig <- syscall.SIGTERM }()
		err = di.Invoke(a.EntryPoint)
	}()
	<-exitSig

	for _, dtor := range a.Dtors {
		if err := di.Invoke(dtor.Fn); err != nil {
			fmt.Fprintf(Stdout, "WARN: %s\n", err.Error())
		}
	}
	return err
}
