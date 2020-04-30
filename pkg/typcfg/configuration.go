package typcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	_ Configurer = &Configuration{}
)

// Configuration is alias from typcore.Configuration with Configurer implementation
type Configuration struct {
	Name string
	Spec interface{}
}

// NewConfiguration return new instance of Configuration
func NewConfiguration(name string, spec interface{}) *Configuration {
	return &Configuration{
		Name: name,
		Spec: spec,
	}
}

// Configurations of configuration instance
func (c *Configuration) Configurations() []*Configuration {
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
