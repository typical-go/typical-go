package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"go.uber.org/dig"
)

var (
	_ctors []*Constructor
	_dtors []*Destructor
)

// Provide constructor
func Provide(ctors ...*Constructor) {
	_ctors = append(_ctors, ctors...)
}

// Destroy destructor
func Destroy(dtors ...*Destructor) {
	_dtors = append(_dtors, dtors...)
}

func launchApp(d *Descriptor) (err error) {
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		_, err = LoadConfig(configFile)
	}

	di := dig.New()
	if err = setDependencies(di, d); err != nil {
		return
	}

	errs := common.GracefulRun(start(di, d), stop(di))
	return errs.Unwrap()
}

func setDependencies(di *dig.Container, d *Descriptor) (err error) {
	if err = di.Provide(func() *Descriptor { return d }); err != nil {
		return
	}
	for _, c := range _ctors {
		if err = di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return
		}
	}
	return
}

func start(di *dig.Container, d *Descriptor) func() error {
	return func() (err error) {
		return di.Invoke(d.EntryPoint)
	}
}

func stop(di *dig.Container) func() error {
	return func() (err error) {
		for _, dtor := range _dtors {
			if err = di.Invoke(dtor.Fn); err != nil {
				return
			}
		}
		return
	}
}
