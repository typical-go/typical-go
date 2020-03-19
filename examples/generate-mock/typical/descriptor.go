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
	Version: "0.0.1",

	App: typapp.
		Create(helloworld.New()),

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(), // standard build module
		),
}
