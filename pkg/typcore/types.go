package typcore

// App is interface of app
type App interface {
	RunApp(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	RunBuildTool(*Context) error
}

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *Context) error
}

// SourceableApp mean the app define its own project sources
type SourceableApp interface {
	ProjectSources() []string
}

// ConfigManager responsible to manage config
type ConfigManager interface {
	Configurations() []*Configuration
	RetrieveConfig(name string) (interface{}, error)
}
