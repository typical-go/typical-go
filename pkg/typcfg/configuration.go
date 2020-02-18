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
type Configurer interface {
	Configure(loader Loader) *Detail
}

// Detail contain detail of config
type Detail struct {
	// Prefix is used by ConfigLoader to retrieve configuration value
	Prefix string

	// Spec is specification of config object
	Spec interface{}

	// Constructor is constructor function to create config object
	Constructor interface{}
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
