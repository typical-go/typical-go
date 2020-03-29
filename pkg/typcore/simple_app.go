package typcore

var (
	_ App = (*SimpleApp)(nil)
)

// SimpleApp is simple implementation of App
type SimpleApp struct {
	fn      func(*Descriptor) error
	sources []string
}

// NewApp return new instance of SimpleApp
func NewApp(fn func(*Descriptor) error, sources ...string) *SimpleApp {
	return &SimpleApp{
		fn:      fn,
		sources: sources,
	}
}

// RunApp to run the simple app
func (a *SimpleApp) RunApp(d *Descriptor) (err error) {
	return a.fn(d)
}

// AppSources return source of the application
func (a *SimpleApp) AppSources() []string {
	return a.sources
}
