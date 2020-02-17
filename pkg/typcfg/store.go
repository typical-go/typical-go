package typcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Store to return config store that contain config informatino
func (c *Configuration) Store() *typcore.ConfigStore {
	store := new(typcore.ConfigStore)
	for _, configurer := range c.configurers {
		prefix, spec, constructor := configurer.Configure(c.loader)
		keys, fieldMap := c.fieldmap(prefix, spec)
		store.Add(&typcore.ConfigBean{
			Constructor: constructor,
			Keys:        keys,
			FieldMap:    fieldMap,
		})
	}

	return store
}

func (c *Configuration) fieldmap(prefix string, spec interface{}) (keys []string, m map[string]*typcore.ConfigField) {
	m = make(map[string]*typcore.ConfigField)
	val := reflect.Indirect(reflect.ValueOf(spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {
			name := fmt.Sprintf("%s_%s", prefix, fieldName(field))
			m[name] = &typcore.ConfigField{
				Name:     name,
				Type:     field.Type.Name(),
				Default:  fieldDefault(field),
				Required: fieldRequired(field),
				Value:    val.Field(i).Interface(),
				IsZero:   val.Field(i).IsZero(),
			}
			keys = append(keys, name)
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
