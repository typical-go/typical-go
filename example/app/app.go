package app

import (
	"fmt"

	"github.com/typical-go/typical-go/example/app/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Module of application
type Module struct {
	Prefix string
}

// New return new instance of application
func New() *Module {
	return &Module{
		Prefix: "APP",
	}
}

// WithPrefix return Module with new prefix
func (m *Module) WithPrefix(prefix string) *Module {
	m.Prefix = prefix
	return m
}

// EntryPoint of application
func (*Module) EntryPoint() interface{} {
	return func(cfg config.Config) {
		fmt.Printf("Hello %s\n", cfg.Hello)
	}
}

// Configure the application
func (m *Module) Configure(loader typcfg.Loader) *typcfg.Detail {
	return &typcfg.Detail{
		Prefix: m.Prefix,
		Spec:   &config.Config{},
		Constructor: func() (cfg config.Config, err error) {
			err = loader.Load(m.Prefix, &cfg)
			return
		},
	}
}
