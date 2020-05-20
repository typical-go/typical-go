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

	Build: typgo.Builds{
		&ReactDemoModule{source: "react-demo"},
		&typgo.StdBuild{},
	},

	Layouts: []string{
		"internal",
	},
}
