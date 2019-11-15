package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typictx"
)

// Context of project
var Context = &typictx.Context{
	Name:        "Typical-Go",
	Description: "Example of typical and scalable RESTful API Server for Go",
	Root:        "github.com/typical-go/typical-go",
	AppModule:   app.Module(),

	Release: typictx.Release{
		Version: "0.9.0",
		Targets: []string{"linux/amd64", "darwin/amd64"},
		Github: &typictx.Github{
			Owner:    "typical-go",
			RepoName: "typical-go",
		},
	},
}
