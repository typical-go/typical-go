package typical

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/stdrls"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Version: typicalgo.Version,
	Package: "github.com/typical-go/typical-go",

	App: typicalgo.New(),

	BuildTool: typbuildtool.New().
		WithRelease(stdrls.New().
			WithPublisher(
				stdrls.GithubPublisher("typical-go", "typical-go"),
			),
		),

	// TODO: remove this when default sources implemented
	Sources: []string{"typicalgo"},
}
