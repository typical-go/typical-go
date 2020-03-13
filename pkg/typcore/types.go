package typcore

// App is interface of app
type App interface {
	Run(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	Run(*Context) error
}

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *Context) error
}

// Sourceable mean the object can return the sources
type Sourceable interface {
	ProjectSources() []string
}

// ConfigManager responsible to manage config
type ConfigManager interface {
	Configurations() []*Configuration
	RetrieveConfig(name string) (interface{}, error)
}
