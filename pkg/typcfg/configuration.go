package typcfg

// Configuration of typical project
type Configuration struct {
	loader      Loader
	configurers []Configurer
}

// Loader responsible to load config
type Loader interface {
	Load(string, interface{}) error
}

// Configurer responsible to create config
// `prefix` is used by ConfigLoader to retrieve configuration value
// `spec` (Specification) is used readme/env file generator. The value of spec will act as local environment value defined in .env file.
// `constructor` (Constructor Function) to be provided in dependecies-injection container
type Configurer interface {
	Configure(loader Loader) (prefix string, spec interface{}, constructor interface{})
}

// New return new instance of Configuration
func New() *Configuration {
	return &Configuration{
		loader: &defaultLoader{},
	}
}

// WithLoader to set loader
func (c *Configuration) WithLoader(loader Loader) *Configuration {
	c.loader = loader
	return c
}

// WithConfigure to set configurer
func (c *Configuration) WithConfigure(configurers ...Configurer) *Configuration {
	c.configurers = append(c.configurers, configurers...)
	return c
}

// Loader of configuration
func (c *Configuration) Loader() Loader {
	return c.loader
}
