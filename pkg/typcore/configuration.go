package typcore

// Configuration is interface of configuration
type Configuration interface {
	Provide() []interface{}
	ConfigMap() ([]string, map[string]ConfigDetail)
	Setup() error
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
