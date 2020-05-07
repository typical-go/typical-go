package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/server"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "server-echo-react",
	Version: "1.0.0",

	App: typcore.NewApp(server.Main),

	BuildTool: &typbuildtool.BuildTool{
		BuildSequences: []interface{}{
			&ReactDemoModule{source: "react-demo"},
			typbuildtool.StandardBuild(),
		},
	},
}
