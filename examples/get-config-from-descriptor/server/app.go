package server

import (
	"fmt"
	"html"
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcore"
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

// Run server
func (a *App) Run(d *typcore.Descriptor) (err error) {
	var spec interface{}
	if spec, err = d.RetrieveConfigSpec(a.ConfigName); err != nil {
		return
	}

	// type assertion to Config type
	cfg := spec.(*Config)

	fmt.Printf("Configuration With Config Store -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
