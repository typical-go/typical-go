package typcore

// App is interface of app
type App interface {
	Run(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	Run(*Context) error
	SetupMe(*Descriptor) error
}

// Configuration is interface of configuration
type Configuration interface {
	Store() *ConfigStore
	Loader() ConfigLoader
	Setup() error
}

// Sourceable mean the object can return the sources
type Sourceable interface {
	ProjectSources() []string
}

// ConfigLoader responsible to load config
type ConfigLoader interface {
	LoadConfig(string, interface{}) error
}

// Configurer responsible to create config
type Configurer interface {
	Configure() *ConfigBean
}
