package typcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	_ Configurer = &Configuration{}
)

// Configuration is alias from typcore.Configuration with Configurer implementation
type Configuration struct {
	*typcore.Configuration
}

// NewConfiguration return new instance of Configuration
func NewConfiguration(name string, spec interface{}) *Configuration {
	return &Configuration{
		Configuration: &typcore.Configuration{
			Name: name,
			Spec: spec,
		},
	}
}

// WithName return Configuration with new name
func (c *Configuration) WithName(name string) *Configuration {
	c.Name = name
	return c
}

// WithSpec return Configuration with new spec
func (c *Configuration) WithSpec(spec interface{}) *Configuration {
	c.Spec = spec
	return c
}

// Configure return the configuration
func (c *Configuration) Configure() []*Configuration {
	return []*Configuration{c}
}

// Fields to retrieve fields from configuration
func (c *Configuration) Fields() (fields []*Field) {
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
