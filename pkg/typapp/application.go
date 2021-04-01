package typapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/errkit"
	"go.uber.org/dig"
)

// defaultExitSigs exit signals
var defaultExitSigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

// Application application
type Application struct {
	StartFn    interface{}
	ShutdownFn interface{}
	ExitSigs   []os.Signal
}

// Run the application
func (a Application) Run() error {
	di := dig.New()
	di.Provide(func() *dig.Container { return di })
	for _, c := range glob {
		if err := di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return err
		}
	}

	exitSigs := a.ExitSigs
	if len(exitSigs) < 1 {
		exitSigs = defaultExitSigs
	}

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, exitSigs...)

	var errs errkit.Errors
	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		errs.Append(di.Invoke(a.StartFn))
	}()
	<-exitCh

	if a.ShutdownFn != nil {
		errs.Append(di.Invoke(a.ShutdownFn))
	}

	return errs.Unwrap()
}
