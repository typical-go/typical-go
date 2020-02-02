package typcore

// Configuration of typical project
type Configuration struct {
	loader      ConfigLoader
	configurers []Configurer
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

// NewConfiguration return new instance of Configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		loader: newDefaultConfigLoader(),
	}
}

// WithLoader to set loader
func (c *Configuration) WithLoader(loader ConfigLoader) *Configuration {
	c.loader = loader
	return c
}

// WithConfigure to set configurer
func (c *Configuration) WithConfigure(configurers ...Configurer) *Configuration {
	c.configurers = append(c.configurers, configurers...)
	return c
}

// Loader of configuration
func (c *Configuration) Loader() ConfigLoader {
	return c.loader
}

// Provide the constructors
func (c *Configuration) Provide() (constructors []interface{}) {
	for _, configurer := range c.configurers {
		_, _, loadFn := configurer.Configure(c.loader)
		constructors = append(constructors, loadFn)
	}
	return
}

// ConfigMap return map of config detail
func (c *Configuration) ConfigMap() (keys []string, configMap ConfigMap) {
	configMap = make(map[string]ConfigDetail)
	for _, configurer := range c.configurers {
		prefix, spec, _ := configurer.Configure(c.loader)
		details := CreateConfigDetails(prefix, spec)
		for _, detail := range details {
			name := detail.Name
			configMap[name] = detail
			keys = append(keys, name)
		}
	}
	return
}
