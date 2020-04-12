package common

import (
	"os"
	"os/signal"
	"syscall"
)

// StartGracefuly to start function with itts graceful stop
func StartGracefuly(startFn func() error, stopFn func() error) (errs Errors) {
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
