package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	EntryPoint: server.Main,

	Configurer: server.Configuration(),

	BuildSequences: []interface{}{
		&typgo.StdBuild{},
	},

	Layouts: []string{
		"server",
	},
}
