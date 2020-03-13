package typcore

// App is interface of app
type App interface {
	Run(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	Run(*Context) error
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
}

// ConfigLoader responsible to load config
type ConfigLoader interface {
	LoadConfig(name string, spec interface{}) error
}

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *PreconditionContext) error
}
