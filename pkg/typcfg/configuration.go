package typcfg

import "github.com/typical-go/typical-go/pkg/typcore"

// Configuration is alias from typcore.Configuration with Configurer implementation
type Configuration struct {
	*typcore.Configuration
}

// NewConfiguration return new instance of Configuration
func NewConfiguration(name string, spec interface{}) *Configuration {
	return &Configuration{
		Configuration: &typcore.Configuration{
			Name: name,
			Spec: spec,
		},
	}
}

// Configure return the configuration
func (c *Configuration) Configure() *Configuration {
	return c
}
