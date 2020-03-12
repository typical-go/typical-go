package typcore

import (
	"github.com/kelseyhightower/envconfig"
)

type defaultLoader struct{}

func (*defaultLoader) LoadConfig(prefix string, v interface{}) error {
	// TODO: deprecate envconfig for consitency between doc, envfile and load config
	return envconfig.Process(prefix, v)
}
