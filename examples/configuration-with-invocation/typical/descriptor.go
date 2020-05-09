package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: server.Main,
		Configurer: server.Configuration(),
	},

	BuildTool: &typcore.BuildTool{
		BuildSequences: []interface{}{
			typcore.StandardBuild(),
		},
		Layouts: []string{
			"server",
		},
	},
}
