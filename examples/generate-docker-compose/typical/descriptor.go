package typical

import (
	"fmt"

	"github.com/typical-go/typical-go/examples/generate-docker-compose/pinger"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	App: typapp.EntryPoint(pinger.Main, "pinger"),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(), // standard build module
		).
		WithUtilities(
			typdocker.Compose(
				redisDockerRecipe,
			),
		),
}

var redisDockerRecipe = &typdocker.Recipe{
	Version: typdocker.V3,
	Services: typdocker.Services{
		"redis": typdocker.Service{
			Image:   "redis:4.0.5-alpine",
			Command: fmt.Sprintf(`redis-server --requirepass "%s"`, "redispass"),
			Ports:   []string{"6379:6379"},
		},
	},
}
