package typcore

import (
	"github.com/kelseyhightower/envconfig"
)

// DefaultConfigLoader is default config loader
type defaultConfigLoader struct{}

func newDefaultConfigLoader() *defaultConfigLoader {
	return &defaultConfigLoader{}
}

func (*defaultConfigLoader) Load(prefix string, v interface{}) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(prefix, v)
}
