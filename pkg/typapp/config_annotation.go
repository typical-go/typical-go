package typapp

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	configTag = "@config"
)

type (
	// ConfigAnnotation handle @config annotation
	// e.g. `@config (prefix: "PREFIX" ctor_name:"CTOR")`
	ConfigAnnotation struct {
		Target  string
		EnvFile bool
	}
	// ConfigAnnotated template
	ConfigAnnotated struct {
		Package string
		Imports []string
		Configs []*Config
	}
	// Config model
	Config struct {
		CtorName string
		Prefix   string
		SpecType string
		Fields   []*Field
	}
	// Field model
	Field struct {
		Key     string
		Default string
	}
)

var _ typannot.Annotator = (*ConfigAnnotation)(nil)

// Annotate config to prepare dependency-injection and env-file
func (m *ConfigAnnotation) Annotate(c *typannot.Context) error {

	configs := m.createConfigs(c)

	if err := m.generate(c, configs); err != nil {
		return err
	}

	target := typgo.EnvFile
	envmap, err := common.CreateEnvMapFromFile(target)
	if err != nil {
		envmap = make(common.EnvMap)
	}

	if m.EnvFile {
		for _, config := range configs {
			for _, field := range config.Fields {
				if _, ok := envmap[field.Key]; !ok {
					envmap[field.Key] = field.Default
				}
			}
		}

		f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := envmap.Save(f); err != nil {
			return err
		}
	}

	if len(envmap) > 0 {
		common.Setenv(envmap)
	}
	return nil
}

func (m *ConfigAnnotation) generate(c *typannot.Context, configs []*Config) error {
	target := m.GetTarget(c)
	if err := common.ExecuteTmplToFile(target, configAnnotTmpl, &ConfigAnnotated{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/kelseyhightower/envconfig",
		),
		Configs: configs,
	}); err != nil {
		return err
	}
	goImports(target)
	return nil
}

func (m *ConfigAnnotation) createConfigs(c *typannot.Context) []*Config {
	var configs []*Config
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckStruct(configTag) {
			prefix := getPrefix(annot)
			var fields []*Field
			for _, field := range annot.Type.(*typannot.StructType).Fields {
				fields = append(fields, &Field{
					Key:     fmt.Sprintf("%s_%s", prefix, getFieldName(field)),
					Default: field.Get("default"),
				})
			}

			configs = append(configs, &Config{
				CtorName: getCtorName(annot),
				Prefix:   prefix,
				SpecType: fmt.Sprintf("%s.%s", annot.Package, annot.Name),
				Fields:   fields,
			})
		}
	}
	return configs
}

// GetTarget get target generation
func (m *ConfigAnnotation) GetTarget(c *typannot.Context) string {
	if m.Target == "" {
		m.Target = fmt.Sprintf("cmd/%s/config_annotated.go", c.BuildSys.Name)
	}
	return m.Target
}

func getCtorName(annot *typannot.Annot) string {
	return annot.TagParam.Get("ctor_name")
}

func getPrefix(annot *typannot.Annot) string {
	prefix := annot.TagParam.Get("prefix")
	if prefix == "" {
		prefix = strings.ToUpper(annot.Name)
	}
	return prefix
}

func getFieldName(field *typannot.Field) string {
	name := field.Get("envconfig")
	if name == "" {
		name = strings.ToUpper(field.Name)
	}
	return name
}
