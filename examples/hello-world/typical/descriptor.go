package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world/helloworld"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	App: typcore.NewApp(helloworld.Main), // the application

	BuildTool: &typbuild.BuildTool{
		BuildSequences: []interface{}{
			typbuild.StandardBuild(), // standard build module
		},
	},
}
