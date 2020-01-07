package typcore

import (
	"github.com/kelseyhightower/envconfig"
)

// DefaultConfigLoader is default config loader
type DefaultConfigLoader struct{}

// NewDefaultConfigLoader return default config loader
func NewDefaultConfigLoader() *DefaultConfigLoader {
	return &DefaultConfigLoader{}
}

// Load configuration
func (*DefaultConfigLoader) Load(prefix string, v interface{}) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(prefix, v)
}
