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

// TypicalConfiguration of typical project
type TypicalConfiguration struct {
	loader    ConfigLoader
	beanNames []string
	beanMap   map[string]*ConfigBean
}

// NewConfiguration return new instance of Configuration
func NewConfiguration() *TypicalConfiguration {
	return &TypicalConfiguration{
		loader:  &defaultLoader{},
		beanMap: make(map[string]*ConfigBean),
	}
}

// WithLoader return TypicalConfiguration with new loader
func (c *TypicalConfiguration) WithLoader(loader ConfigLoader) *TypicalConfiguration {
	c.loader = loader
	return c
}

// WithConfigurer return TypicalConfiguratiton with new configurers
func (c *TypicalConfiguration) WithConfigurer(configurers ...Configurer) *TypicalConfiguration {
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
func (c *TypicalConfiguration) Loader() ConfigLoader {
	return c.loader
}

// Setup the configuration to be ready to use for the app and build-tool
func (c *TypicalConfiguration) Setup() (err error) {
	var (
		f *os.File
	)
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		log.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()
		if err = c.Write(f); err != nil {
			return
		}
	}
	// TODO: load env
	return
}

// Write typical configuration
func (c *TypicalConfiguration) Write(w io.Writer) (err error) {
	for _, bean := range c.Beans() {
		for _, field := range bean.Fields() {
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
func (c *TypicalConfiguration) Put(bean *ConfigBean) {
	name := bean.Name
	if _, exist := c.beanMap[name]; exist {
		panic(fmt.Sprintf("Can't put '%s' to config store", name))
	}
	c.beanNames = append(c.beanNames, name)
	c.beanMap[name] = bean
}

// Get bean from config store
func (c *TypicalConfiguration) Get(name string) *ConfigBean {
	return c.beanMap[name]
}

// Beans return array of bean
func (c *TypicalConfiguration) Beans() (beans []*ConfigBean) {
	for _, name := range c.beanNames {
		beans = append(beans, c.beanMap[name])
	}
	return
}
