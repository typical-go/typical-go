package typical

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typgo.Descriptor{

	Name:    "typical-go",
	Version: "0.9.50",

	EntryPoint: wrapper.Main,

	Layouts: []string{
		"wrapper",
		"pkg",
	},

	BuildSequences: []interface{}{
		typgo.StandardBuild(),
		&typgo.Github{
			Owner:    "typical-go",
			RepoName: "typical-go",

			PublishSetting: typgo.PublishSetting{
				ExcludeMessage: typgo.ExcludePrefix("merge", "bump", "revision", "generate", "wip"),
			},
		},
	},

	Utility: typgo.NewUtility(taskTestExample), // Test all the examples

}
