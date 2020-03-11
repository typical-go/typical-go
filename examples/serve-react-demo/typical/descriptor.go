package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/server"
	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "server-echo-react",
	Version: "0.0.1",

	App: serverApp,

	BuildTool: typbuildtool.New().
		WithBuilder(typbuildtool.NewBuilder().
			Before(
				exor.NewCommand("npm", "run", "build").WithDir("react-demo"),
			),
		),
}

// Modules
var (
	serverApp = server.New()
)
