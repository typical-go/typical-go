package typical

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
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

	BuildTool: &typcore.BuildTool{
		Layouts: []string{
			"wrapper",
			"pkg",
		},
		BuildSequences: []interface{}{
			typcore.StandardBuild(),
			typcore.Github("typical-go", "typical-go"),
		},
		Utility: typcore.NewUtility(taskTestExample), // Test all the examples
	},
}
