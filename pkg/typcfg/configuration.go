package typcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Configuration represent the configuration
type Configuration struct {
	Prefix string
	Spec   interface{}
}

// Field contain field information of spec
type Field struct {
	Name     string
	Type     string
	Default  string
	Required bool
}

// Configure return config itself
func (c Configuration) Configure() Configuration {
	return c
}

// ConfigFields return list of field information
func (c Configuration) ConfigFields() (infos []Field) {
	val := reflect.Indirect(reflect.ValueOf(c.Spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {
			infos = append(infos, Field{
				Name:     fmt.Sprintf("%s_%s", c.Prefix, fieldName(field)),
				Type:     field.Type.Name(),
				Default:  fieldDefault(field),
				Required: fieldRequired(field),
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

func fieldIgnored(field reflect.StructField) (required bool) {
	if v, ok := field.Tag.Lookup("ignored"); ok {
		required, _ = strconv.ParseBool(v)
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
