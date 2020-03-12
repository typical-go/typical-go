package server

import (
	"fmt"
	"html"
	"net/http"

	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server/config"
	"github.com/typical-go/typical-go/pkg/typcore"
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
func (a *App) Configure() *typcore.ConfigBean {
	return &typcore.ConfigBean{
		Name: a.ConfigName,
		Spec: &config.Config{},
	}
}

// Run server
func (a *App) Run(d *typcore.Descriptor) (err error) {
	cfgBean := d.Configuration.Store().Get(a.ConfigName)
	loader := d.Configuration.Loader()

	if err = loader.LoadConfig(cfgBean.Name, cfgBean.Spec); err != nil {
		return
	}

	cfg := cfgBean.Spec.(*config.Config)

	fmt.Printf("Configuration With Config Store -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
