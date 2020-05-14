package common

import (
	"os"
	"os/signal"
	"syscall"
)

type (
	// Fn is simple function
	Fn func() error
)

// GracefulRun with start and end function
func GracefulRun(startFn Fn, stopFn Fn) (errs Errors) {
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
