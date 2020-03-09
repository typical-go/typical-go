package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/examples/configuration-with-invocation/helloworld/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typdep"
)

// Module of application
type Module struct {
	ConfigName string
}

// New return new instance of application
func New() *Module {
	return &Module{
		ConfigName: "APP",
	}
}

// WithConfigPrefix return Module with new config prefix
func (m *Module) WithConfigPrefix(name string) *Module {
	m.ConfigName = name
	return m
}

// EntryPoint of application
func (*Module) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(
		func(cfg config.Config) {
			fmt.Printf("Hello %s\n", cfg.Hello)
		})
}

// Configure the application
func (m *Module) Configure(loader typcfg.Loader) *typcfg.Configuration {
	return &typcfg.Configuration{
		Name: m.ConfigName,
		Spec: &config.Config{},
		Constructor: typdep.NewConstructor(
			func() (cfg config.Config, err error) {
				err = loader.Load(m.ConfigName, &cfg)
				return
			}),
	}
}
