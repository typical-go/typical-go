package typical

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/pinger"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	App: typcore.Run(pinger.Main),

	BuildTool: &typbuild.BuildTool{
		BuildSequences: []interface{}{
			typbuild.StandardBuild(), // standard build module
		},
		Utilities: []typbuild.Utility{
			typdocker.Compose(redisRecipe),
		},
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
