package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/examples/configuration-with-invocation/helloworld/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typdep"
)

// App of hello world
type App struct {
	ConfigName string
}

// New return new instance of application
func New() *App {
	return &App{
		ConfigName: "APP",
	}
}

// WithConfigPrefix return Module with new config prefix
func (a *App) WithConfigPrefix(name string) *App {
	a.ConfigName = name
	return a
}

// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(
		func(cfg config.Config) {
			fmt.Printf("Hello %s\n", cfg.Hello)
		})
}

// Configure the application
func (a *App) Configure(loader typcfg.Loader) *typcfg.Configuration {
	return &typcfg.Configuration{
		Name: a.ConfigName,
		Spec: &config.Config{},
		Constructor: typdep.NewConstructor(
			func() (cfg config.Config, err error) {
				err = loader.Load(a.ConfigName, &cfg)
				return
			}),
	}
}
