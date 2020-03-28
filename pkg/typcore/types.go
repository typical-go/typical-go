package typcore

// App is interface of app
type App interface {
	RunApp(*Descriptor) error
	AppSources() []string
}

// BuildTool interface
type BuildTool interface {
	RunBuildTool(*Context) error
}

// ConfigManager responsible to manage config
type ConfigManager interface {
	Configurations() []*Configuration
	RetrieveConfig(name string) (interface{}, error)
}

// Configuration is detail of config
type Configuration struct {
	Name string
	Spec interface{}
}
