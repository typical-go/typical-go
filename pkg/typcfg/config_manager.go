package typcfg

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

const (
	defaultDotEnv = ".env"
)

var (
	_ typcore.ConfigManager       = (*ConfigManager)(nil)
	_ typbuildtool.Preconditioner = (*ConfigManager)(nil)
)

// ConfigManager of typical project
type ConfigManager struct {
	loader      Loader
	configurers []Configurer
}

// Configures to create new instance of ConfigManager with set of configurers
func Configures(configurers ...Configurer) *ConfigManager {
	return &ConfigManager{
		configurers: configurers,
		loader:      &defaultLoader{},
	}
}

// WithLoader return TypicalConfiguration with new loader
func (m *ConfigManager) WithLoader(loader Loader) *ConfigManager {
	m.loader = loader
	return m
}

// Loader of configuration
func (m *ConfigManager) Loader() Loader {
	return m.loader
}

// Precondition to use config manager
func (m *ConfigManager) Precondition(c *typbuildtool.BuildContext) (err error) {
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
func (m *ConfigManager) Write(w io.Writer) (err error) {
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
func (m *ConfigManager) RetrieveConfig(name string) (interface{}, error) {
	cfg := m.Get(name)
	spec := cfg.Spec
	if err := m.LoadConfig(name, spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// Configurations return array of configuration
func (m *ConfigManager) Configurations() (cfgs []*typcore.Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configure().Configuration)
	}
	return
}

// Get the configuration
func (m *ConfigManager) Get(name string) *typcore.Configuration {
	for _, cfg := range m.Configurations() {
		if cfg.Name == name {
			return cfg
		}
	}
	return nil
}

// LoadConfig to load the config
func (m *ConfigManager) LoadConfig(name string, spec interface{}) error {
	if m.loader != nil {
		return m.loader.LoadConfig(name, spec)
	}
	return fmt.Errorf("ConfigLoader is missing")
}

func loadEnvFile(c *typbuildtool.BuildContext) (err error) {
	source := ".env"

	envMap, err := ReadFile(source)
	if err != nil {
		return
	}

	if len(envMap) > 0 {
		c.Infof("Load environments %s", source)
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
