package typical

import (
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: typcore.Version,

	App: app.New(),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
			typbuildtool.Github("typical-go", "typical-go"),
		).
		Utilities(
			typbuildtool.NewUtility(taskTestExample), // Test all the examples
		),
}
