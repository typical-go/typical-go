package typcore

// Configuration is interface of configuration
// TODO: refactor to minimum function requirement
type Configuration interface {
	Provide() []interface{}
	Loader() ConfigLoader
	ConfigMap() ([]string, map[string]ConfigDetail)
	Setup() error
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

// ConfigDetail is detail field of config
type ConfigDetail struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}

// ConfigDetailsBy to return config details from map bases-on keys
func ConfigDetailsBy(c map[string]ConfigDetail, keys ...string) (details []ConfigDetail) {
	for _, key := range keys {
		if detail, ok := c[key]; ok {
			details = append(details, detail)
		}
	}
	return
}
