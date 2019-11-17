package typcfg

import (
	"github.com/kelseyhightower/envconfig"
)

// Loader responsible to load config
type Loader interface {
	Load(Configuration) error
}

// DefaultLoader return default config loader
func DefaultLoader() Loader {
	return defaultLoader{}
}

type defaultLoader struct{}

// Load configuration
func (defaultLoader) Load(c Configuration) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(c.Prefix, c.Spec)
}
