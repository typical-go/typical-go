package typcore

import (
	"github.com/kelseyhightower/envconfig"
)

// ConfigLoader responsible to load config
type ConfigLoader interface {
	Load(string, interface{}) error
}

// DefaultConfigLoader return default config loader
func DefaultConfigLoader() ConfigLoader {
	return &defaultConfigLoader{}
}

type defaultConfigLoader struct{}

// Load configuration
func (defaultConfigLoader) Load(prefix string, v interface{}) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(prefix, v)
}
