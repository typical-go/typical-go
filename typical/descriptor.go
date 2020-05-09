package typical

import (
	"github.com/typical-go/typical-go/pkg/github"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typgo.Descriptor{

	Name:    "typical-go",
	Version: "0.9.50",

	App: &typapp.App{
		EntryPoint: wrapper.Main,
	},

	BuildTool: &typgo.BuildTool{
		Layouts: []string{
			"wrapper",
			"pkg",
		},
		BuildSequences: []interface{}{
			typgo.StandardBuild(),
			github.Github("typical-go", "typical-go"),
		},
		Utility: typgo.NewUtility(taskTestExample), // Test all the examples
	},
}
