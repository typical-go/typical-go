package app

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/example/app/config"
)

// Module of application
type Module struct{}

// New return new instance of application
func New() *Module {
	return &Module{}
}

// EntryPoint of application
func (*Module) EntryPoint() interface{} {
	return func(cfg config.Config) {
		fmt.Printf("Hello %s\n", cfg.Hello)
	}
}

// Configure the application
func (*Module) Configure(loader typcfg.Loader) (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func() (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
