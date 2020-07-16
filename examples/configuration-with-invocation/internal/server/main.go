package server

import (
	"fmt"
	"net/http"
)

type (
	// Config of app
	Config struct {
		Address string `default:":8080" required:"true"`
	}
)

// Main function to run server
func Main(cfg *Config) error {
	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, &handler{})
}
