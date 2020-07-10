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

var _ Compiler = (*ConfigManager)(nil)

// Compile to prepare dependency-injection and env-file
func (m *ConfigManager) Compile(c *Context) error {
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

	cfgGenerated := fmt.Sprintf("%s/%s/cfg_generated.go", CmdFolder, c.Descriptor.Name)
	if err := writeGoSource(&typtmpl.CfgGenerated{
		Package:  "main",
		Imports:  c.Imports,
		CfgCtors: cfgs,
	}, cfgGenerated); err != nil {
		return err
	}

	if !m.SkipEnvFile {
		if err := WriteConfig(ConfigFile, m.Configs); err != nil {
			return err
		}
	}

	return nil
}
