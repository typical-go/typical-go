package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/stdrls"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Version: app.Version,
	Package: "github.com/typical-go/typical-go",

	App: typapp.New(application),

	Build: typbuild.New().
		WithRelease(stdrls.New().
			WithPublisher(
				stdrls.GithubPublisher("typical-go", "typical-go"),
			),
		),
}
