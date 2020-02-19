package typical

import (
	"github.com/typical-go/typical-go/example/app"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	application = app.New()

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

		// Package of the project (MANDATORY)
		// It should be same with go.mod file
		Package: "github.com/typical-go/typical-go/example",

		// App of the project (MANDATORY)
		App: typapp.New(application),

		// BuildTool of the project (MANDATORY)
		BuildTool: typbuildtool.New(),

		// Configuration of the project (OPTIONAL)
		Configuration: typcfg.New().
			WithConfigure(
				application,
			),
	}
)
