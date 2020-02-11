package typcore

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Configuration is interface of configuration
type Configuration interface {
	Provide() []interface{}
	Loader() ConfigLoader
	ConfigMap() (keys []string, configMap ConfigMap)
	Setup() error
}

// ConfigLoader responsible to load config
type ConfigLoader interface {
	Load(string, interface{}) error
}

// Configurer responsible to create config
// `Prefix` is used by ConfigLoader to retrieve configuration value
// `Spec` (Specification) is used readme/env file generator. The value of spec will act as local environment value defined in .env file.
// `LoadFn` (Load Function) is required to provide in dependecies-injection container
type Configurer interface {
	Configure(loader ConfigLoader) (prefix string, spec interface{}, loadFn interface{})
}

// ConfigDetail is detail of config
type ConfigDetail struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}

// ConfigMap is map of config detail
type ConfigMap map[string]ConfigDetail

// CreateConfigDetails is mapping of config field
func CreateConfigDetails(prefix string, spec interface{}) (details ConfigDetails) {
	val := reflect.Indirect(reflect.ValueOf(spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {
			details = append(details, ConfigDetail{
				Name:     fmt.Sprintf("%s_%s", prefix, fieldName(field)),
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

// ConfigDetails is slice of ConfigDetail
type ConfigDetails []ConfigDetail

// ValueBy to return values by key
func (c *ConfigMap) ValueBy(keys ...string) (details ConfigDetails) {
	for _, key := range keys {
		if detail, ok := (*c)[key]; ok {
			details = append(details, detail)
		}
	}
	return
}

// Append ConfigDetail
func (c *ConfigDetails) Append(detail ConfigDetail) *ConfigDetails {
	*c = append(*c, detail)
	return c
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
