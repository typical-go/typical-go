package server

import (
	"fmt"
	"html"
	"net/http"
)

const (
	// ConfigName for server lookup key in config manaager
	ConfigName = "SERVER"
)

// Config of app
type Config struct {
	Address string `default:":8080" required:"true"`
}

// Main function to run server
func Main(cfg *Config) error {
	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, &handler{})
}

type handler struct{}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
