package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/server"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "server-echo-react",
	Version: "0.0.1",

	App: serverApp,

	BuildTool: typbuildtool.New().WithModules(
		NewReactDemoModule(),
		typbuildtool.NewModule(),
	),
}

// Modules
var (
	serverApp = server.New()
)
