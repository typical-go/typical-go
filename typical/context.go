package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of project
var Context = &typctx.Context{
	Name:      "Typical-Go",
	Version:   app.Version,
	Package:   "github.com/typical-go/typical-go",
	AppModule: app.Module(),
	Releaser: &typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
		Publishers: []typrls.Publisher{
			&typrls.Github{Owner: "typical-go", RepoName: "typical-go"},
		},
	},
	ConfigLoader: typcfg.DefaultLoader(),
}
