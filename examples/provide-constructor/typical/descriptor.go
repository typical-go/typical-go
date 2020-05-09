package typical

import (
	"github.com/typical-go/typical-go/examples/provide-constructor/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "provide-constructor",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: helloworld.Main2,
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(),
		},
		Layouts: []string{
			"helloworld",
		},
	},
}
