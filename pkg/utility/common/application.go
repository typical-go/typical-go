package common

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Application with start and graceful stop function
type Application struct {
	StartFn func() error
	StopFn  func() error
}

// Run the application
func (r *Application) Run() (errs coll.Errors) {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		if err := r.StartFn(); err != nil {
			errs.Append(err)
		}
	}()
	<-gracefulStop
	if r.StopFn == nil {
		if err := r.StopFn(); err != nil {
			errs.Append(err)
		}
	}
	return
}
