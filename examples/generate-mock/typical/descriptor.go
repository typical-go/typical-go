package typical

import (
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-mock",
	Version: "1.0.0",

	App: typapp.AppModule(helloworld.New()),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(), // standard build module
		),
}
