package typapp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"
)

type (
	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

var _ctors []*Constructor

// Provide constructor
func Provide(name string, fn interface{}) {
	_ctors = append(_ctors, &Constructor{Name: name, Fn: fn})
}

// Reset constructor
func Reset() {
	_ctors = make([]*Constructor, 0)
}

// Run the entry points
func Run(startFn interface{}, shutdownFns ...interface{}) error {
	di := dig.New()
	for _, c := range _ctors {
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
		err = di.Invoke(startFn)
	}()
	<-exitSig

	for _, shutdownFn := range shutdownFns {
		if err1 := di.Invoke(shutdownFn); err1 != nil {
			fmt.Printf("shutdown: %s", err1.Error())
		}
	}

	return err
}
