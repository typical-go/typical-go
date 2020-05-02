package server

import (
	"fmt"
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcfg"
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

// Configuration of server
func Configuration() *typcfg.Configuration {
	return &typcfg.Configuration{
		Name: ConfigName,
		Spec: &Config{},
	}
}
