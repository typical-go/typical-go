package typical

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: typcore.Version,

	App: typicalgo.Create(),

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(),
			typbuildtool.Github("typical-go", "typical-go"),
		).
		WithCommanders(
			typbuildtool.CreateCommander(taskTestExample), // Test all the examples
		),
}
