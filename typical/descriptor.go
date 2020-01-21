package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var (
	application = app.New()

	// Descriptor of project
	Descriptor = typcore.ProjectDescriptor{
		Name:    "Typical-Go",
		Version: app.Version,
		Package: "github.com/typical-go/typical-go",

		App: typcore.NewApp().
			WithCommand(
				application,
			),

		Releaser: typrls.New().
			WithPublisher(
				typrls.GithubPublisher("typical-go", "typical-go"),
			),
	}
)
