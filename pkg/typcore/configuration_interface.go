package typcore

// ConfigurationInterface is interface of configuration
type ConfigurationInterface interface {
	Provider
	Loader() ConfigLoader
	ConfigMap() (keys []string, configMap ConfigMap)
}
