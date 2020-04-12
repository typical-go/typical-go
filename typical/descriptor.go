package typical

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: "0.9.49",

	App: wrapper.New(),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
			typbuildtool.Github("typical-go", "typical-go"),
		).
		Utilities(
			typbuildtool.NewUtility(taskTestExample), // Test all the examples
		),
}
