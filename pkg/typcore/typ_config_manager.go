package typcore

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	defaultDotEnv = ".env"
)

// TypConfigManager of typical project
type TypConfigManager struct {
	loader    ConfigLoader
	beanNames []string
	beanMap   map[string]*Configuration
}

// NewConfigManager return new instance of Configuration
func NewConfigManager() *TypConfigManager {
	return &TypConfigManager{
		loader:  &defaultLoader{},
		beanMap: make(map[string]*Configuration),
	}
}

// WithLoader return TypicalConfiguration with new loader
func (c *TypConfigManager) WithLoader(loader ConfigLoader) *TypConfigManager {
	c.loader = loader
	return c
}

// WithConfigurer return TypicalConfiguratiton with new configurers
func (c *TypConfigManager) WithConfigurer(configurers ...Configurer) *TypConfigManager {
	for _, configurer := range configurers {
		cfg := configurer.Configure()
		if cfg == nil {
			panic("Configure return nil detail")
		}
		c.Put(cfg)
	}
	return c
}

// Loader of configuration
func (c *TypConfigManager) Loader() ConfigLoader {
	return c.loader
}

// Setup the configuration to be ready to use for the app and build-tool
func (c *TypConfigManager) Setup() (err error) {
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()

		log.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if err = c.Write(f); err != nil {
			return
		}
	}
	// TODO: load env
	return
}

// Write typical configuration
func (c *TypConfigManager) Write(w io.Writer) (err error) {
	for _, cfg := range c.Configurations() {
		for _, field := range cfg.Fields() {
			var v interface{}
			if field.IsZero {
				v = field.Default
			} else {
				v = field.Value
			}
			if _, err = fmt.Fprintf(w, "%s=%v\n", field.Name, v); err != nil {
				return
			}
		}
	}
	return
}

// Put bean to config store
func (c *TypConfigManager) Put(bean *Configuration) {
	name := bean.Name
	if _, exist := c.beanMap[name]; exist {
		panic(fmt.Sprintf("Can't put '%s' to config store", name))
	}
	c.beanNames = append(c.beanNames, name)
	c.beanMap[name] = bean
}

// GetConfig to get configuration
func (c *TypConfigManager) GetConfig(name string) *Configuration {
	return c.beanMap[name]
}

// Configurations return array of configuration
func (c *TypConfigManager) Configurations() (beans []*Configuration) {
	for _, name := range c.beanNames {
		beans = append(beans, c.beanMap[name])
	}
	return
}
