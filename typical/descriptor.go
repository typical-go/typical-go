package typical

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: "0.9.50",

	App: &typapp.App{
		EntryPoint: wrapper.Main,
	},

	BuildTool: &typbuild.BuildTool{
		Layouts: []string{
			"wrapper",
			"pkg",
		},
		BuildSequences: []interface{}{
			typbuild.StandardBuild(),
			typbuild.Github("typical-go", "typical-go"),
		},
		Utility: typbuild.NewUtility(taskTestExample), // Test all the examples
	},
}
