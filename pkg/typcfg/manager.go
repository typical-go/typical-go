package typcfg

import (
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	_ typcore.ConfigManager       = (*ConfigManager)(nil)
	_ typbuildtool.Preconditioner = (*ConfigManager)(nil)
)

// ConfigManager of typical project
type ConfigManager struct {
	configurers []Configurer
	source      string
}

// Configures to create new instance of ConfigManager with set of configurers
func Configures(configurers ...Configurer) *ConfigManager {
	return &ConfigManager{
		source:      ".env",
		configurers: configurers,
	}
}

// Precondition to use config manager
func (m *ConfigManager) Precondition(c *typbuildtool.BuildContext) (err error) {
	var (
		envMap map[string]string
		b      strings.Builder
	)

	if _, err = os.Stat(m.source); os.IsNotExist(err) {
		c.Infof("Generate new project environment at '%s'", m.source)
		if err = Write(m, m.source); err != nil {
			return
		}
	}

	if envMap, err = Read(m.source); err != nil {
		return
	}

	if len(envMap) > 0 {
		c.Infof("Load environments %s", m.source)
		for key, value := range envMap {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString("+" + key + " ")
		}
		c.Info(b.String())
	}
	return
}

// RetrieveConfig to get configuration spec
func (m *ConfigManager) RetrieveConfig(name string) (interface{}, error) {
	cfg := m.Get(name)
	spec := cfg.Spec
	if err := envconfig.Process(name, spec); err != nil {
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
