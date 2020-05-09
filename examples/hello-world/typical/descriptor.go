package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world/helloworld"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	App: &typgo.App{
		EntryPoint: helloworld.Main,
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(), // standard build module
		},
	},
}
