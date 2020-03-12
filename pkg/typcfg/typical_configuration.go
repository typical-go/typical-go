package typcfg

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
)

const (
	defaultDotEnv = ".env"
)

// TypicalConfiguration of typical project
type TypicalConfiguration struct {
	loader      Loader
	configurers []Configurer
}

// New return new instance of Configuration
func New() *TypicalConfiguration {
	return &TypicalConfiguration{
		loader: &defaultLoader{},
	}
}

// WithLoader to set loader
func (c *TypicalConfiguration) WithLoader(loader Loader) *TypicalConfiguration {
	c.loader = loader
	return c
}

// AppendConfigurer to append configurer
func (c *TypicalConfiguration) AppendConfigurer(configurers ...Configurer) *TypicalConfiguration {
	c.configurers = append(c.configurers, configurers...)
	return c
}

// Store to return config store that contain config informatino
func (c *TypicalConfiguration) Store() *typcore.ConfigStore {
	store := typcore.NewConfigStore()
	for _, configurer := range c.configurers {
		cfg := configurer.Configure()
		if cfg == nil {
			panic("Configure return nil detail")
		}
		store.Put(cfg)
	}

	return store
}

// Loader of configuration
func (c *TypicalConfiguration) Loader() Loader {
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
	store := c.Store()
	for _, field := range store.Fields() {
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
	return
}
