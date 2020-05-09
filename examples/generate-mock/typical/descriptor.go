package typical

import (
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-mock",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: helloworld.Main,
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(), // standard build module
		},
		Utility: typmock.Utility(),
		Layouts: []string{
			"helloworld",
		},
	},
}
