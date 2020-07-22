package server

import (
	"fmt"
	"html"
	"net/http"
)

type (
	// Config of app
	// @config
	Config struct {
		Address string `default:":8080" required:"true"`
	}
	handler struct{}
)

// Main function to run server
func Main(cfg *Config) error {
	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, &handler{})
}

//
// handler
//

var _ http.Handler = (*handler)(nil)

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
