package common

import (
	"os"
	"os/signal"
	"syscall"
)

// Application with start and graceful stop function
type Application struct {
	startFn func() error
	stopFn  func() error
}

// NewApplication to return new instance of Application
func NewApplication(startFn func() error) *Application {
	return &Application{
		startFn: startFn,
	}
}

// WithStopFn to set stop function
func (a *Application) WithStopFn(stopFn func() error) *Application {
	a.stopFn = stopFn
	return a
}

// Run the application
func (a *Application) Run() (errs Errors) {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		if err := a.startFn(); err != nil {
			// NOTE: if startFn got error, it should still execute stopFn
			errs.Append(err)
		}
	}()
	<-gracefulStop
	if a.stopFn != nil {
		if err := a.stopFn(); err != nil {
			errs.Append(err)
		}
	}
	return
}
