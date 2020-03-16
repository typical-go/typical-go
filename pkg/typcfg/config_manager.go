package typcfg

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

const (
	defaultDotEnv = ".env"
	configKey     = "CONFIG"
)

// TypConfigManager of typical project
type TypConfigManager struct {
	loader      Loader
	configurers []Configurer
}

// New instance of Configuration
func New() *TypConfigManager {
	return &TypConfigManager{
		loader: &defaultLoader{},
	}
}

// WithLoader return TypicalConfiguration with new loader
func (m *TypConfigManager) WithLoader(loader Loader) *TypConfigManager {
	m.loader = loader
	return m
}

// WithConfigurers return TypicalConfiguratiton with new configurers
func (m *TypConfigManager) WithConfigurers(configurers ...Configurer) *TypConfigManager {
	m.configurers = configurers
	return m
}

// Loader of configuration
func (m *TypConfigManager) Loader() Loader {
	return m.loader
}

// Precondition to use config manager
func (m *TypConfigManager) Precondition(c *typbuildtool.Context) (err error) {
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()

		c.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if err = m.Write(f); err != nil {
			return
		}
	}
	if err = loadEnvFile(c); err != nil {
		return
	}
	return
}

// Write typical configuration
func (m *TypConfigManager) Write(w io.Writer) (err error) {
	for _, cfg := range m.Configurations() {
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

// RetrieveConfig to get configuration spec
func (m *TypConfigManager) RetrieveConfig(name string) (interface{}, error) {
	cfgdef := m.Get(name)
	spec := cfgdef.Spec()
	if err := m.LoadConfig(cfgdef.Name(), spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// Configurations return array of configuration
func (m *TypConfigManager) Configurations() (cfgs []*typcore.Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configure())
	}
	return
}

// Get the configuration
func (m *TypConfigManager) Get(name string) *typcore.Configuration {
	for _, cfg := range m.Configurations() {
		if cfg.Name() == name {
			return cfg
		}
	}
	return nil
}

// LoadConfig to load the config
func (m *TypConfigManager) LoadConfig(name string, spec interface{}) error {
	if m.loader != nil {
		return m.loader.LoadConfig(name, spec)
	}
	return fmt.Errorf("ConfigLoader is missing")
}

func loadEnvFile(c *typbuildtool.Context) (err error) {
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
		c.Infof("Load environments %s", configSource)
		var b strings.Builder
		for key, value := range envMap {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString("+")
			b.WriteString(key)
			b.WriteString(" ")
		}
		c.Info(b.String())
	}
	return
}
