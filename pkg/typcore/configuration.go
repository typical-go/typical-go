package typcore

// Configuration is detail of config
type Configuration struct {
	name string
	spec interface{}
}

// NewConfiguration return new instance of Configuration
func NewConfiguration(name string, spec interface{}) *Configuration {
	return &Configuration{
		name: name,
		spec: spec,
	}
}

// Name of configuration
func (c *Configuration) Name() string {
	return c.name
}

// Spec of configuration
func (c *Configuration) Spec() interface{} {
	return c.spec
}
