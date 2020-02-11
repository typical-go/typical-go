package typcfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// ConfigMap return map of config detail
func (c *Configuration) ConfigMap() (keys []string, m map[string]typcore.ConfigDetail) {
	m = make(map[string]typcore.ConfigDetail)
	for _, configurer := range c.configurers {
		prefix, spec, _ := configurer.Configure(c.loader)
		val := reflect.Indirect(reflect.ValueOf(spec))
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if !fieldIgnored(field) {
				name := fmt.Sprintf("%s_%s", prefix, fieldName(field))
				m[name] = typcore.ConfigDetail{
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
