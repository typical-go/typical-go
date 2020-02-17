package typcore

// ConfigStore contain information of config
type ConfigStore struct {
	beans []*ConfigBean
}

// ConfigBean is detail of config
type ConfigBean struct {
	Constructor interface{}
	Keys        []string
	FieldMap    map[string]*ConfigField
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

// Add bean to config store
func (c *ConfigStore) Add(bean *ConfigBean) {
	c.beans = append(c.beans, bean)
}

// Provide list of functino
func (c *ConfigStore) Provide() (constructors []interface{}) {
	for _, bean := range c.beans {
		constructors = append(constructors, bean.Constructor)
	}
	return
}

// Keys return all config key
func (c *ConfigStore) Keys() (keys []string) {
	for _, bean := range c.beans {
		keys = append(keys, bean.Keys...)
	}
	return
}

// FieldMap return field map
func (c *ConfigStore) FieldMap() (m map[string]*ConfigField) {
	m = make(map[string]*ConfigField)
	for _, bean := range c.beans {
		for k, v := range bean.FieldMap {
			m[k] = v
		}
	}
	return
}

// Fields return array of field by keys
func (c *ConfigStore) Fields(keys ...string) (fields []*ConfigField) {
	m := c.FieldMap()
	for _, key := range keys {
		if detail, ok := m[key]; ok {
			fields = append(fields, detail)
		}
	}
	return
}
