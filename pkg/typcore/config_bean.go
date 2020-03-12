package typcore

import "github.com/typical-go/typical-go/pkg/typdep"

// ConfigBean is detail of config
type ConfigBean struct {
	name        string
	fields      []*ConfigField
	constructor *typdep.Constructor
}

// ConfigField is detail field of config
type ConfigField struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}

// NewConfigBean return new instance of ConfigBean
func NewConfigBean(name string, fields []*ConfigField, constructor *typdep.Constructor) *ConfigBean {
	return &ConfigBean{
		name:        name,
		fields:      fields,
		constructor: constructor,
	}
}

// Name of Config Bean
func (c *ConfigBean) Name() string {
	return c.name
}

// Fields of Config Bean
func (c *ConfigBean) Fields() []*ConfigField {
	return c.fields
}

// Constructor of Config Bean
func (c *ConfigBean) Constructor() *typdep.Constructor {
	return c.constructor
}
