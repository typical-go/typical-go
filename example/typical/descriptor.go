package typical

import (
	"github.com/typical-go/typical-go/example/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	hello = helloworld.New()

	// Descriptor of sample
	Descriptor = typcore.Descriptor{

		// Name of the project (OPTIONAL)
		// It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		// Name: "example",

		// Description of the project (OPTIONAL)
		Description: "Example of typical and scalable RESTful API Server for Go",

		// Version of the project (MANDATORY)
		Version: "0.0.1",

		// ModulePackage of the project (MANDATORY)
		// Normally it should not be set as it will retrieve from `go.mod` file or project path after the $GOPATH
		ModulePackage: "github.com/typical-go/typical-go/example",

		// App of the project (MANDATORY)
		App: typapp.New(hello),

		// BuildTool of the project (MANDATORY)
		BuildTool: typbuildtool.New(),

		// Configuration of the project (OPTIONAL)
		Configuration: typcfg.New().
			WithConfigure(
				hello,
			),
	}
)
