package typapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/errkit"
)

// ExitSigs exit signals
var ExitSigs = []syscall.Signal{syscall.SIGTERM, syscall.SIGINT}

// Run the application
func Run(startFn, shutdownFn interface{}) error {
	di, err := Container()
	if err != nil {
		return err
	}

	exitCn := make(chan os.Signal)
	for _, s := range ExitSigs {
		signal.Notify(exitCn, s)
	}

	var errs errkit.Errors
	go func() {
		defer func() { exitCn <- syscall.SIGTERM }()
		errs.Append(di.Invoke(startFn))
	}()
	<-exitCn

	if shutdownFn != nil {
		errs.Append(di.Invoke(shutdownFn))
	}

	return errs.Unwrap()
}
