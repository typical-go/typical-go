package typical

import (
	"github.com/typical-go/typical-go/examples/generate-mock/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typmock"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "generate-mock",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: helloworld.Main,
	},

	BuildTool: &typbuild.BuildTool{
		BuildSequences: []interface{}{
			typbuild.StandardBuild(), // standard build module
		},
		Utility: typmock.Utility(),
		Layouts: []string{
			"helloworld",
		},
	},
}
