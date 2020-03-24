package typcfg

import "github.com/typical-go/typical-go/pkg/typcore"

// Configurer responsible to create config
type Configurer interface {
	Configure() *typcore.Configuration
}

// ConfigLoader responsible to load config
type Loader interface {
	LoadConfig(name string, spec interface{}) error
}

// Field of config
type Field struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}
