package typcfg

import "github.com/typical-go/typical-go/pkg/typdep"

// Loader responsible to load config
type Loader interface {
	Load(string, interface{}) error
}

// Configurer responsible to create config
type Configurer interface {
	Configure(loader Loader) *Configuration
}

// Configuration contain detail of config
type Configuration struct {
	// Name of configuration
	Name string

	// Spec is specification of config object. Spec used to generate the .env file
	Spec interface{}

	// Constructor is constructor function to create config object
	Constructor *typdep.Constructor
}
