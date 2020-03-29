package typcfg

import "github.com/typical-go/typical-go/pkg/typcore"

var (
	_ Configurer = &Configuration{}
)

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

// WithName return Configuration with new name
func (c *Configuration) WithName(name string) *Configuration {
	c.Name = name
	return c
}

// WithSpec return Configuration with new spec
func (c *Configuration) WithSpec(spec interface{}) *Configuration {
	c.Spec = spec
	return c
}

// Configure return the configuration
func (c *Configuration) Configure() *Configuration {
	return c
}
