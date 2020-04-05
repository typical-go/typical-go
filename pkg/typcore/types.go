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
