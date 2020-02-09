package typcore

import "io"

// Configuration is interface of configuration
type Configuration interface {
	Provider
	Loader() ConfigLoader
	ConfigMap() (keys []string, configMap ConfigMap)
	Write(io.Writer) error
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
