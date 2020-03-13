package typcfg

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
)

const (
	defaultDotEnv = ".env"
	configKey     = "CONFIG"
)

// TypConfigManager of typical project
type TypConfigManager struct {
	loader      typcore.ConfigLoader
	configurers []Configurer
}

// Configurer responsible to create config
type Configurer interface {
	Configure() *typcore.Configuration
}

// New instance of Configuration
func New() *TypConfigManager {
	return &TypConfigManager{
		loader: &defaultLoader{},
	}
}

// WithLoader return TypicalConfiguration with new loader
func (c *TypConfigManager) WithLoader(loader typcore.ConfigLoader) *TypConfigManager {
	c.loader = loader
	return c
}

// WithConfigurer return TypicalConfiguratiton with new configurers
func (c *TypConfigManager) WithConfigurer(configurers ...Configurer) *TypConfigManager {
	c.configurers = configurers
	return c
}

// Loader of configuration
func (c *TypConfigManager) Loader() typcore.ConfigLoader {
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
	if err = loadEnvFile(); err != nil {
		return
	}
	return
}

// Write typical configuration
func (c *TypConfigManager) Write(w io.Writer) (err error) {
	for _, cfg := range c.Configurations() {
		for _, field := range RetrieveFields(cfg) {
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

// RetrieveConfigSpec to get configuration spec
func (c *TypConfigManager) RetrieveConfigSpec(name string) (interface{}, error) {
	cfgdef := c.Get(name)
	spec := cfgdef.Spec()
	if err := c.LoadConfig(cfgdef.Name(), spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// Configurations return array of configuration
func (c *TypConfigManager) Configurations() (cfgs []*typcore.Configuration) {
	for _, configurer := range c.configurers {
		cfgs = append(cfgs, configurer.Configure())
	}
	return
}

// Get the configuration
func (c *TypConfigManager) Get(name string) *typcore.Configuration {
	for _, cfg := range c.Configurations() {
		if cfg.Name() == name {
			return cfg
		}
	}
	return nil
}

// LoadConfig to load the config
func (c *TypConfigManager) LoadConfig(name string, spec interface{}) error {
	if c.loader != nil {
		return c.loader.LoadConfig(name, spec)
	}
	return fmt.Errorf("ConfigLoader is missing")
}

func loadEnvFile() (err error) {
	// TODO: don't use godotenv for flexibility
	configSource := os.Getenv(configKey)
	var configs []string
	var envMap map[string]string
	if configSource == "" {
		envMap, _ = godotenv.Read()
	} else {
		configs = strings.Split(configSource, ",")
		if envMap, err = godotenv.Read(configs...); err != nil {
			return
		}
	}

	if len(envMap) > 0 {
		log.Infof("Load environments %s", configSource)
		var b strings.Builder
		for key, value := range envMap {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString("+")
			b.WriteString(key)
			b.WriteString(" ")
		}
		log.Info(b.String())
	}
	return
}
