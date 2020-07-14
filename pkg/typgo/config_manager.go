package typgo

import (
	"fmt"
	"reflect"

	"github.com/typical-go/typical-go/pkg/typtmpl"
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

var _ Action = (*ConfigManager)(nil)

// Execute config-manager to prepare dependency-injection and env-file
func (m *ConfigManager) Execute(c *Context) error {
	var cfgs []*typtmpl.CfgCtor
	for _, cfg := range m.Configs {
		specType := reflect.TypeOf(cfg.Spec).String()
		cfgs = append(cfgs, &typtmpl.CfgCtor{
			Name:      cfg.Ctor,
			Prefix:    cfg.Name,
			SpecType:  specType,
			SpecType2: specType[1:],
		})
	}

	if err := writeGoSource(
		fmt.Sprintf("cmd/%s/config_annotated.go", c.Descriptor.Name),
		&typtmpl.ConfigAnnotated{
			Package:  "main",
			Imports:  c.Imports,
			CfgCtors: cfgs,
		},
	); err != nil {
		return err
	}

	if !m.SkipEnvFile {
		if err := WriteConfig(ConfigFile, m.Configs); err != nil {
			return err
		}
	}

	return nil
}
