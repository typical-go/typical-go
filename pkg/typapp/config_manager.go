package typapp

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type (
	// ConfigManager manage the configs
	ConfigManager struct {
		Target  string
		EnvFile bool
		Configs []*Config
	}
	// Config model
	Config struct {
		Ctor   string
		Prefix string
		Spec   interface{}
	}
	// Field of config
	Field struct {
		Name     string
		Type     string
		Default  string
		Value    interface{}
		IsZero   bool
		Required bool
	}
)

var _ typast.Annotator = (*ConfigManager)(nil)

// Annotate config to prepare dependency-injection and env-file
func (m *ConfigManager) Annotate(c *typast.Context) error {
	if err := m.execute(c); err != nil {
		return err
	}
	if m.EnvFile {
		if err := m.Save(typgo.EnvFile); err != nil {
			return err
		}
	}
	return nil
}

func (m *ConfigManager) execute(c *typast.Context) error {
	var cfgs []*typtmpl.CfgCtor
	for _, cfg := range m.Configs {
		specType := reflect.TypeOf(cfg.Spec).String()
		cfgs = append(cfgs, &typtmpl.CfgCtor{
			Name:      cfg.Ctor,
			Prefix:    cfg.Prefix,
			SpecType:  specType,
			SpecType2: specType[1:],
		})
	}

	return WriteGoSource(
		m.GetTarget(c),
		&typtmpl.ConfigAnnotated{Package: "main", Imports: c.Imports, CfgCtors: cfgs},
	)
}

// GetTarget get target generation
func (m *ConfigManager) GetTarget(c *typast.Context) string {
	if m.Target == "" {
		m.Target = fmt.Sprintf("cmd/%s/config_annotated.go", c.Descriptor.Name)
	}
	return m.Target
}

// Save to file
func (m *ConfigManager) Save(target string) error {
	envmap, err := common.CreateEnvMapFromFile(target)
	if err != nil {
		envmap = make(common.EnvMap)
	}

	for _, field := range m.Fields() {
		if _, ok := envmap[field.Name]; !ok {
			envmap[field.Name] = fmt.Sprintf("%v", field.GetValue())
		}
	}

	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	return envmap.Save(f)
}

// Fields of config
func (m *ConfigManager) Fields() []*Field {
	var fields []*Field
	for _, cfg := range m.Configs {
		for _, field := range CreateFields(cfg) {
			fields = append(fields, field)
		}
	}
	return fields
}

// CreateFields to retrieve fields from configuration
func CreateFields(c *Config) (fields []*Field) {
	val := reflect.Indirect(reflect.ValueOf(c.Spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {

			fields = append(fields, &Field{
				Name:     fmt.Sprintf("%s_%s", c.Prefix, fieldName(field)),
				Type:     field.Type.Name(),
				Default:  fieldDefault(field),
				Required: fieldRequired(field),
				Value:    val.Field(i).Interface(),
				IsZero:   val.Field(i).IsZero(),
			})
		}
	}
	return
}

func fieldRequired(field reflect.StructField) (required bool) {
	if v, ok := field.Tag.Lookup("required"); ok {
		required, _ = strconv.ParseBool(v)
	}
	return
}

func fieldIgnored(field reflect.StructField) (ignored bool) {
	if v, ok := field.Tag.Lookup("ignored"); ok {
		ignored, _ = strconv.ParseBool(v)
	}
	return
}

func fieldDefault(field reflect.StructField) string {
	return field.Tag.Get("default")
}

func fieldName(field reflect.StructField) (name string) {
	name = strings.ToUpper(field.Name)
	if v, ok := field.Tag.Lookup("envconfig"); ok {
		name = v
	}
	return
}

// GetValue to get value or default value if no value
func (f *Field) GetValue() interface{} {
	if f.IsZero {
		return f.Default
	}
	return f.Value
}
