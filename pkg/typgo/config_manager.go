package typgo

import (
	"reflect"

	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

type (
	// ConfigManager manage the configs
	ConfigManager struct {
		SkipEnvFile bool
		Configs     []*Configuration
	}

	// Configuration is alias from typgo.Configuration with Configurer implementation
	Configuration struct {
		Ctor string
		Name string
		Spec interface{}
	}
)

var _ Prebuilder = (*ConfigManager)(nil)

// Prebuild to prepare dependency-injection and env-file
func (m *ConfigManager) Prebuild(c *Context) error {
	for _, cfg := range m.Configs {
		specType := reflect.TypeOf(cfg.Spec).String()
		c.Precond.CfgCtors = append(c.Precond.CfgCtors, &typtmpl.CfgCtor{
			Name:      cfg.Ctor,
			Prefix:    cfg.Name,
			SpecType:  specType,
			SpecType2: specType[1:],
		})
	}

	if !m.SkipEnvFile {
		if err := WriteConfig(typvar.ConfigFile, m.Configs); err != nil {
			return err
		}
	}

	return nil
}
