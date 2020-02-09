package typcfg

import (
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Configuration of typical project
type Configuration struct {
	loader      typcore.ConfigLoader
	configurers []typcore.Configurer
}

// New return new instance of Configuration
func New() *Configuration {
	return &Configuration{
		loader: newDefaultConfigLoader(),
	}
}

// WithLoader to set loader
func (c *Configuration) WithLoader(loader typcore.ConfigLoader) *Configuration {
	c.loader = loader
	return c
}

// WithConfigure to set configurer
func (c *Configuration) WithConfigure(configurers ...typcore.Configurer) *Configuration {
	c.configurers = append(c.configurers, configurers...)
	return c
}

// Loader of configuration
func (c *Configuration) Loader() typcore.ConfigLoader {
	return c.loader
}

// Provide the constructors
func (c *Configuration) Provide() (constructors []interface{}) {
	for _, configurer := range c.configurers {
		_, _, loadFn := configurer.Configure(c.loader)
		constructors = append(constructors, loadFn)
	}
	return
}
