package typical

import (
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typcore.Descriptor{

	Name:    "typical-go",
	Version: "0.9.50",

	App: typcore.Run(wrapper.Main),

	BuildTool: &typbuild.BuildTool{
		BuildSequences: []interface{}{
			typbuild.StandardBuild(),
			typbuild.Github("typical-go", "typical-go"),
		},
		Utilities: []typbuild.Utility{
			typbuild.NewUtility(taskTestExample), // Test all the examples
		},
	},

	Layouts: []string{
		"wrapper",
		"pkg",
	},
}
