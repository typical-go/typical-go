package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/internal/server"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	EntryPoint: server.Main,
	Layouts: []string{
		"internal",
	},

	Prebuild: &typgo.ConfigManager{
		Configs: []*typgo.Configuration{
			{Name: "SERVER", Spec: &server.Config{}},
		},
	},

	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Clean:   &typgo.StdClean{},
}
