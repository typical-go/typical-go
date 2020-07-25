package app

import (
	"fmt"
	"html"
	"net/http"
)

type (
	// ServerCfg configuration
	// @cfg (prefix:"SERVER")
	ServerCfg struct {
		Address string `envconfig:"ADDRESS" default:":8080" required:"true"`
	}
	handler struct{}
)

// Main function to run server
func Main(cfg *ServerCfg) error {
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
