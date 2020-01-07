package typcore

// Configuration of typical project
type Configuration struct {
	loader      ConfigLoader
	configurers []Configurer
}

// Configurer responsible to create config
// `Prefix` is used by ConfigLoader to retrieve configuration value
// `Spec` (Specification) is used readme/env file generator. The value of spec will act as local environment value defined in .env file.
// `LoadFn` (Load Function) is required to provide in dependecies-injection container
type Configurer interface {
	Configure() (prefix string, spec interface{}, loadFn interface{})
}

// NewConfiguration return new instance of Configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		loader: DefaultConfigLoader(),
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

// Provide the constructors
func (c *Configuration) Provide() (constructors []interface{}) {
	constructors = append(constructors, func() ConfigLoader {
		return c.loader
	})
	for _, configurer := range c.configurers {
		_, _, loadFn := configurer.Configure()
		constructors = append(constructors, loadFn)
	}
	return
}
