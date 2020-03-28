package typical

import (
	"github.com/typical-go/typical-go/examples/simple-additional-task/helloworld"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "simple-additional-task",
	Version: "1.0.0",

	App: typcore.NewApp(helloworld.Main),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
		).
		WithUtilities(
			typbuildtool.NewUtility(taskPrintContext), // Add custom task
		),
}
