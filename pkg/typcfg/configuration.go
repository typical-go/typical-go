package typcfg

import (
	"fmt"
	"io"

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

// ConfigMap return map of config detail
func (c *Configuration) ConfigMap() (keys []string, configMap typcore.ConfigMap) {
	configMap = make(map[string]typcore.ConfigDetail)
	for _, configurer := range c.configurers {
		prefix, spec, _ := configurer.Configure(c.loader)
		details := typcore.CreateConfigDetails(prefix, spec)
		for _, detail := range details {
			name := detail.Name
			configMap[name] = detail
			keys = append(keys, name)
		}
	}
	return
}

func (c *Configuration) Write(w io.Writer) (err error) {
	keys, configMap := c.ConfigMap()
	for _, key := range keys {
		var (
			v         interface{}
			cfgDetail = configMap[key]
		)
		if cfgDetail.IsZero {
			v = cfgDetail.Default
		} else {
			v = cfgDetail.Value
		}
		if _, err = fmt.Fprintf(w, "%s=%v\n", cfgDetail.Name, v); err != nil {
			return
		}
	}
	return
}
