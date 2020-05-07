package typical

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: "0.9.50",

	App: wrapper.New(),

	BuildTool: &typbuildtool.BuildTool{
		BuildSequences: []interface{}{
			typbuildtool.StandardBuild(),
			typbuildtool.Github("typical-go", "typical-go"),
		},
		Utilities: []typbuildtool.Utility{
			typbuildtool.NewUtility(taskTestExample), // Test all the examples
		},
	},

	Layouts: []string{
		"wrapper",
		"pkg",
	},
}
