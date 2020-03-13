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

	App: typicalgo.New(),

	BuildTool: typbuildtool.New().
		AppendCommanders(
			typbuildtool.NewCommander(taskTestExample), // Test all the examples
		).
		WithPublishers(
			typbuildtool.NewGithub("typical-go", "typical-go"),
		),
}
