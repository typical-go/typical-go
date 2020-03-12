package typcfg

import (
	"github.com/kelseyhightower/envconfig"
)

type defaultLoader struct{}

// NewDefaultLoader return new instance of default loader
func NewDefaultLoader() Loader {
	return &defaultLoader{}
}

func (*defaultLoader) Load(prefix string, v interface{}) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(prefix, v)
}
