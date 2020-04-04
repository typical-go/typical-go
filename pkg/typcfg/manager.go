package typcfg

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	_ typcore.ConfigManager       = (*ConfigManager)(nil)
	_ Configurer                  = (*ConfigManager)(nil)
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
	c.Infof("Generate new project environment at '%s'", m.source)

	if err = Write(m.source, m); err != nil {
		return
	}

	if _, err = Load(m.source); err != nil {
		return
	}

	return
}

// RetrieveConfig to get configuration spec
func (m *ConfigManager) RetrieveConfig(name string) (interface{}, error) {
	cfg := m.Get(name)
	spec := cfg.Spec
	if err := Process(name, spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// Configurations return array of configuration
// TODO: remove this
func (m *ConfigManager) Configurations() (cfgs []*typcore.Configuration) {
	for _, configurer := range m.configurers {
		for _, configurations := range configurer.Configure() {
			cfgs = append(cfgs, configurations.Configuration)
		}
	}
	return
}

// Configure the application or buildtool
func (m *ConfigManager) Configure() (cfgs []*Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configure()...)
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
