package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/internal/server"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "server-echo-react",
	Version: "1.0.0",

	EntryPoint: server.Main,
	Layouts: []string{
		"internal",
	},

	Compile: typgo.Compiles{
		&ReactDemoModule{source: "react-demo"},
		&typgo.StdCompile{},
	},

	Run: &typgo.StdRun{},
}
