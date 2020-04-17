package typical

import (
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typmock"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-mock",
	Version: "1.0.0",

	App: typapp.EntryPoint(helloworld.Main, "helloworld"),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(), // standard build module
		).
		Utilities(
			typmock.Utility(),
		),
}
