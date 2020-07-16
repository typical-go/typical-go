package typapp

var (
	// Name of application. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typapp.Name=PROJECT-NAME`
	Name string
	// Version of applicatoin. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typapp.Version=PROJECT-NAME`
	Version string
)
