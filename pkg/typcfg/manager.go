package typcfg

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

var (
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

// Configurations of manager
func (m *ConfigManager) Configurations() (cfgs []*Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configurations()...)
	}
	return
}
