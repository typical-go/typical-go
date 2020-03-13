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

// Sourceable mean the object can return the sources
type Sourceable interface {
	ProjectSources() []string
}

// ConfigManager responsible to manage config
type ConfigManager interface {
	ConfigLoader
	Configurations() []*Configuration
	RetrieveConfigSpec(name string) (interface{}, error)

	Setup() error // TODO: remove this
}

// ConfigLoader responsible to load config
type ConfigLoader interface {
	LoadConfig(name string, spec interface{}) error
}
