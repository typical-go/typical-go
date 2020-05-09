package typical

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/pinger"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: pinger.Main,
	},

	BuildTool: &typcore.BuildTool{
		BuildSequences: []interface{}{
			typcore.StandardBuild(), // standard build module
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
