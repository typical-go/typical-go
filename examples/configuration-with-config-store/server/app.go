package server

import (
	"fmt"
	"html"
	"net/http"

	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server/config"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
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

// Configure the application
func (a *App) Configure(loader typcfg.Loader) *typcfg.Configuration {
	return &typcfg.Configuration{
		Name: a.ConfigName,
		Spec: &config.Config{},
		Constructor: typdep.NewConstructor(func() (cfg config.Config, err error) {
			err = loader.Load(a.ConfigName, &cfg)
			return
		}),
	}
}

// Run server
func (a *App) Run(d *typcore.Descriptor) (err error) {
	cfgBean := d.Configuration.Store().Get(a.ConfigName)
	fn := cfgBean.Constructor().Fn().(func() (cfg config.Config, err error))
	var cfg config.Config
	if cfg, err = fn(); err != nil {
		return
	}

	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
