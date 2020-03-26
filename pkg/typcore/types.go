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

// Logger responsible to log any useful information
type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
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
