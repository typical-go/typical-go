package typcfg

import (
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Loader responsible to load config
type Loader interface {
	Load(string, interface{}) error
}

// Configurer responsible to create config
type Configurer interface {
	Configure(loader Loader) *typcore.ConfigBean
}
