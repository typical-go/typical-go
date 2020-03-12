package typcore

import (
	"fmt"
)

// ConfigStore contain information of config
type ConfigStore struct {
	beanNames []string
	beanMap   map[string]*ConfigBean
}

// NewConfigStore return new instance of ConfigStore
func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		beanMap: make(map[string]*ConfigBean),
	}
}

// Put bean to config store
func (c *ConfigStore) Put(bean *ConfigBean) {
	name := bean.Name
	if _, exist := c.beanMap[name]; exist {
		panic(fmt.Sprintf("Can't put '%s' to config store", name))
	}
	c.beanNames = append(c.beanNames, name)
	c.beanMap[name] = bean
}

// Get bean from config store
func (c *ConfigStore) Get(name string) *ConfigBean {
	return c.beanMap[name]
}

// Beans return array of bean
func (c *ConfigStore) Beans() (beans []*ConfigBean) {
	for _, name := range c.beanNames {
		beans = append(beans, c.beanMap[name])
	}
	return
}

// Fields return field map
func (c *ConfigStore) Fields() (fields []*ConfigField) {
	for _, name := range c.beanNames {
		fields = append(fields, c.beanMap[name].Fields()...)
	}
	return
}
