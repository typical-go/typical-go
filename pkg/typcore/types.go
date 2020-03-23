package typcore

// AppLauncher responsible to launch the application
type AppLauncher interface {
	LaunchApp() error
}

// BuildToolLauncher responsible to launch the build-tool
type BuildToolLauncher interface {
	LaunchBuildTool() error
}

// App is interface of app
type App interface {
	RunApp(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	RunBuildTool(*Context) error
}

// SourceableApp mean the app define its own project sources
type SourceableApp interface {
	ProjectSources() []string
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
