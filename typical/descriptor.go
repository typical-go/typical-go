package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Descriptor of project
var Descriptor = &typcore.ProjectDescriptor{
	Name:      "Typical-Go",
	Version:   app.Version,
	Package:   "github.com/typical-go/typical-go",
	AppModule: app.Module(),
	Releaser: typrls.New().WithPublisher(
		typrls.GithubPublisher("typical-go", "typical-go"),
	),
}
