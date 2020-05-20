package server

import "github.com/typical-go/typical-go/pkg/typgo"

const (
	// ConfigName for server lookup key in config manaager
	ConfigName = "SERVER"
)

// Config of app
type Config struct {
	Address string `default:":8080" required:"true"`
}

// Configuration of server
func Configuration() *typgo.Configuration {
	return &typgo.Configuration{
		Name: ConfigName,
		Spec: &Config{},
	}
}
