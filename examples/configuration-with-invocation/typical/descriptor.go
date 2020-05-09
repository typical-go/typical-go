package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: server.Main,
		Configurer: server.Configuration(),
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(),
		},
		Layouts: []string{
			"server",
		},
	},
}
