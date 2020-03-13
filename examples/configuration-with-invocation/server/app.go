package server

import (
	"fmt"
	"html"
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
)

// App of hello world
type App struct {
	ConfigName string
}

// Config of app
type Config struct {
	Address string `default:":8080" required:"true"`
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
func (a *App) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(a.ConfigName, &Config{})
}

// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(a.start)
}

func (a *App) start(cfg *Config) error {
	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
