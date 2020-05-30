package typical

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/internal/pinger"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	EntryPoint: pinger.Main,
	Layouts: []string{
		"internal",
	},

	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},

	Utility: typdocker.Compose(
		redisRecipe,
	),
}
