package typical

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/pinger"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	App: &typgo.App{
		EntryPoint: pinger.Main,
	},

	BuildTool: &typgo.BuildTool{
		BuildSequences: []interface{}{
			typgo.StandardBuild(), // standard build module
		},
		Utility: typdocker.Compose(redisRecipe),
	},
}

var redisRecipe = &typdocker.Recipe{
	Version: typdocker.V3,
	Services: typdocker.Services{
		"redis": typdocker.Service{
			Image: "redis:4.0.5-alpine",
			Ports: []string{"6379:6379"},
		},
		"webdis": typdocker.Service{
			Image: "anapsix/webdis",
			Ports: []string{"7379:7379"},
		},
	},
}
