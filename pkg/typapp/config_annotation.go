package typapp

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	configTag = "@config"
)

type (
	// ConfigAnnotation manage the configs
	ConfigAnnotation struct {
		Target  string
		EnvFile bool
	}
)

var _ typannot.Annotator = (*ConfigAnnotation)(nil)

// Annotate config to prepare dependency-injection and env-file
func (m *ConfigAnnotation) Annotate(c *typannot.Context) error {
	var cfgs []*typtmpl.CfgCtor
	var annots []*typannot.Annot
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckStruct(configTag) {
			annots = append(annots, annot)
			cfgs = append(cfgs, &typtmpl.CfgCtor{
				Name:     annot.TagAttrs.Get("name"),
				Prefix:   strcase.ToScreamingSnake(annot.Name),
				SpecType: fmt.Sprintf("%s.%s", annot.Package, annot.Name),
			})
		}
	}

	target := m.GetTarget(c)
	if err := WriteGoSource(target, &typtmpl.ConfigAnnotated{
		Package:  "main",
		Imports:  c.Imports,
		CfgCtors: cfgs,
	}); err != nil {
		return err
	}
	if m.EnvFile {
		if err := SaveEnvFile(typgo.EnvFile, annots); err != nil {
			return err
		}
	}
	return nil
}

// GetTarget get target generation
func (m *ConfigAnnotation) GetTarget(c *typannot.Context) string {
	if m.Target == "" {
		m.Target = fmt.Sprintf("cmd/%s/config_annotated.go", c.BuildSys.Name)
	}
	return m.Target
}

// SaveEnvFile save env file
func SaveEnvFile(target string, annots []*typannot.Annot) error {
	envmap, err := common.CreateEnvMapFromFile(target)
	if err != nil {
		envmap = make(common.EnvMap)
	}

	for _, annot := range annots {
		prefix := strcase.ToScreamingSnake(annot.Name)
		if structType, ok := annot.Type.(*typannot.StructType); ok {
			for _, field := range structType.Fields {
				name := field.Get("envconfig")
				if name == "" {
					name = strcase.ToScreamingSnake(field.Name)
				}
				key := fmt.Sprintf("%s_%s", prefix, name)
				if _, ok := envmap[key]; !ok {
					envmap[key] = field.Get("default")
				}
			}
		}
	}

	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	return envmap.Save(f)
}
