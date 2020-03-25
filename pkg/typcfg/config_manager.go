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

// TypicalConfigManager of typical project
type TypicalConfigManager struct {
	loader      Loader
	configurers []Configurer
}

// Configures to create new instance of TypicalConfigManager with set of configurers
func Configures(configurers ...Configurer) *TypicalConfigManager {
	return &TypicalConfigManager{
		configurers: configurers,
		loader:      &defaultLoader{},
	}
}

// WithLoader return TypicalConfiguration with new loader
func (m *TypicalConfigManager) WithLoader(loader Loader) *TypicalConfigManager {
	m.loader = loader
	return m
}

// Loader of configuration
func (m *TypicalConfigManager) Loader() Loader {
	return m.loader
}

// Precondition to use config manager
func (m *TypicalConfigManager) Precondition(c *typbuildtool.BuildContext) (err error) {
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
func (m *TypicalConfigManager) Write(w io.Writer) (err error) {
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
func (m *TypicalConfigManager) RetrieveConfig(name string) (interface{}, error) {
	cfg := m.Get(name)
	spec := cfg.Spec
	if err := m.LoadConfig(name, spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// Configurations return array of configuration
func (m *TypicalConfigManager) Configurations() (cfgs []*typcore.Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configure().Configuration)
	}
	return
}

// Get the configuration
func (m *TypicalConfigManager) Get(name string) *typcore.Configuration {
	for _, cfg := range m.Configurations() {
		if cfg.Name == name {
			return cfg
		}
	}
	return nil
}

// LoadConfig to load the config
func (m *TypicalConfigManager) LoadConfig(name string, spec interface{}) error {
	if m.loader != nil {
		return m.loader.LoadConfig(name, spec)
	}
	return fmt.Errorf("ConfigLoader is missing")
}

func loadEnvFile(c *typbuildtool.BuildContext) (err error) {
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
