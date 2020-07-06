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
	exitSig := make(chan os.Signal)
	signal.Notify(exitSig, syscall.SIGTERM)
	signal.Notify(exitSig, syscall.SIGINT)
	go func() {
		defer func() {
			exitSig <- syscall.SIGTERM
		}()
		if err := startFn(); err != nil {
			// NOTE: if startFn got error, it should still execute stopFn
			errs.Append(err)
		}
	}()
	<-exitSig
	if stopFn != nil {
		if err := stopFn(); err != nil {
			errs.Append(err)
		}
	}
	return
}
