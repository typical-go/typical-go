package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of project
var Context = &typctx.Context{
	Name:        "Typical-Go",
	Description: "Example of typical and scalable RESTful API Server for Go",
	Package:     "github.com/typical-go/typical-go",
	AppModule:   typctx.NewAppModule(app.Start),

	Releaser: typrls.Releaser{
		Version:   "0.9.0",
		Targets:   []typrls.Target{"linux/amd64", "darwin/amd64"},
		Publisher: &typrls.Github{Owner: "typical-go", RepoName: "typical-go"},
	},
}
