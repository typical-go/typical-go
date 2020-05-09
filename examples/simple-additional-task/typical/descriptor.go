package typical

import (
	"github.com/typical-go/typical-go/examples/simple-additional-task/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "simple-additional-task",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: helloworld.Main,
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(),
		},
		Utility: typgo.NewUtility(taskPrintContext), // Add custom task
	},
}
