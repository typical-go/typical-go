package typapp

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type (
	// ConfigManager manage the configs
	ConfigManager struct {
		Target  string
		EnvFile bool
		Configs []*Configuration
	}
	// Configuration is alias from typapp.Configuration with Configurer implementation
	Configuration struct {
		Ctor string
		Name string
		Spec interface{}
	}
)

var _ typgo.Action = (*ConfigManager)(nil)

// Execute config-manager to prepare dependency-injection and env-file
func (m *ConfigManager) Execute(c *typgo.Context) error {
	if err := m.execute(c); err != nil {
		return err
	}
	if m.EnvFile {
		if err := WriteConfig(typgo.EnvFile, m.Configs); err != nil {
			return err
		}
	}
	return nil
}

func (m *ConfigManager) execute(c *typgo.Context) error {
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

	return typgo.WriteGoSource(
		m.GetTarget(c),
		&typtmpl.ConfigAnnotated{Package: "main", Imports: c.Imports, CfgCtors: cfgs},
	)
}

// GetTarget get target generation
func (m *ConfigManager) GetTarget(c *typgo.Context) string {
	if m.Target == "" {
		m.Target = fmt.Sprintf("cmd/%s/config_annotated.go", c.Descriptor.Name)
	}
	return m.Target
}

// WriteConfig to write configuration to file
func WriteConfig(dest string, configs []*Configuration) (err error) {
	var fields []*Field
	for _, cfg := range configs {
		for _, field := range CreateFields(cfg) {
			fields = append(fields, field)
		}
	}

	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	hasNewLine, err := hasLastNewLine(f)
	if err != nil {
		return
	}

	if !hasNewLine {
		fmt.Fprintln(f)
	}

	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return
	}

	m := typgo.ReadConfig(f)
	for _, field := range fields {
		if _, ok := m[field.Name]; !ok {
			fmt.Fprintf(f, "%s=%v\n", field.Name, field.GetValue())
		}
	}

	return
}

func hasLastNewLine(f *os.File) (has bool, err error) {
	stat, err := f.Stat()
	if err != nil {
		return
	}

	if stat.Size() <= 0 {
		return true, nil
	}

	if _, err = f.Seek(-1, io.SeekEnd); err != nil {
		return
	}

	char := make([]byte, 1)
	if _, err = f.Read(char); err != nil {
		return
	}

	return (char[0] == '\n'), nil
}

// CreateFields to retrieve fields from configuration
func CreateFields(c *Configuration) (fields []*Field) {
	val := reflect.Indirect(reflect.ValueOf(c.Spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {
			name := fmt.Sprintf("%s_%s", c.Name, fieldName(field))
			fields = append(fields, &Field{
				Name:     name,
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

// Field of config
type Field struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}

// GetValue to get value or default value if no value
func (f *Field) GetValue() interface{} {
	if f.IsZero {
		return f.Default
	}
	return f.Value
}
