package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/typical-go/typical-go/pkg/typirelease"
)

// Context of project
var Context = &typictx.Context{
	Name:        "Typical-Go",
	Description: "Example of typical and scalable RESTful API Server for Go",
	Package:     "github.com/typical-go/typical-go",
	AppModule:   typictx.NewAppModule(app.Start),

	Releaser: typirelease.Releaser{
		Version:   "0.9.0",
		Targets:   []string{"linux/amd64", "darwin/amd64"},
		Publisher: &typirelease.Github{Owner: "typical-go", RepoName: "typical-go"},
	},
}
