package server

import "github.com/typical-go/typical-go/pkg/typapp"

const (
	// ConfigName for server lookup key in config manaager
	ConfigName = "SERVER"
)

// Config of app
type Config struct {
	Address string `default:":8080" required:"true"`
}

// Configuration of server
func Configuration() *typapp.Configuration {
	return &typapp.Configuration{
		Name: ConfigName,
		Spec: &Config{},
	}
}
