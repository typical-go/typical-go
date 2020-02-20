package typcore

// App is interface of app
type App interface {
	Run(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	Run(*TypicalContext) error
}

// Configuration is interface of configuration
type Configuration interface {
	Store() *ConfigStore
	Setup() error
}

// Sourceable mean the object can return the sources
type Sourceable interface {
	ProjectSources() []string
}
