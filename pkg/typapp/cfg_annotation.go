package typapp

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// CfgAnnotation handle @cfg annotation
	// e.g. `@cfg (prefix: "PREFIX" ctor_name:"CTOR")`
	CfgAnnotation struct {
		// Target code generation
		Target string
		// If true then create and load envfile
		DotEnv bool
	}
	// CfgTmplData template
	CfgTmplData struct {
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

var _ typannot.Annotator = (*CfgAnnotation)(nil)

// Annotate config to prepare dependency-injection and env-file
func (m *CfgAnnotation) Annotate(c *typannot.Context) error {
	configs := m.createConfigs(c)

	target := m.getTarget(c)
	data := &CfgTmplData{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg, "github.com/kelseyhightower/envconfig"),
		Configs: configs,
	}
	if err := common.ExecuteTmplToFile(target, configAnnotTmpl, data); err != nil {
		return err
	}
	goImports(target)
	if m.DotEnv {
		if err := CreateAndLoadDotEnv(".env", configs); err != nil {
			return err
		}
	}

	return nil
}

// CreateAndLoadDotEnv to create and load envfile
func CreateAndLoadDotEnv(envfile string, configs []*Config) error {
	envmap, err := common.CreateEnvMapFromFile(envfile)
	if err != nil {
		envmap = make(common.EnvMap)
	}

	var updatedKeys []string
	for _, config := range configs {
		for _, field := range config.Fields {
			if _, ok := envmap[field.Key]; !ok {
				updatedKeys = append(updatedKeys, "+"+field.Key)
				envmap[field.Key] = field.Default
			}
		}
	}
	if len(updatedKeys) > 0 {
		color.New(color.FgGreen).Fprint(Stdout, "UPDATE_ENV")
		fmt.Fprintln(Stdout, ": "+strings.Join(updatedKeys, " "))
	}

	if err := envmap.SaveToFile(envfile); err != nil {
		return err
	}

	return common.Setenv(envmap)
}

func (m *CfgAnnotation) createConfigs(c *typannot.Context) []*Config {
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
func (m *CfgAnnotation) getTarget(c *typannot.Context) string {
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
