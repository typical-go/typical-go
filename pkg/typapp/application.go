package typapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/errkit"
)

// defaultExitSigs exit signals
var defaultExitSigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

// Application application
type Application struct {
	StartFn  interface{}
	StopFn   interface{}
	ExitSigs []os.Signal
}

// StartService start the service with gracefully stop
func StartService(startFn, stopFn interface{}, exitSigs ...os.Signal) error {

	if len(exitSigs) < 1 {
		exitSigs = defaultExitSigs
	}

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, exitSigs...)

	var errs errkit.Errors
	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		errs.Append(Invoke(startFn))
	}()
	<-exitCh

	if stopFn != nil {
		errs.Append(Invoke(stopFn))
	}

	return errs.Unwrap()
}
