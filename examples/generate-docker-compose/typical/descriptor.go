package typical

import (
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	Layouts: []string{"internal"},

	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Clean:   &typgo.StdClean{},

	Utility: &typdocker.Utility{
		Version: typdocker.V3,
		Composers: []typdocker.Composer{
			redisRecipe,
		},
	},
}
