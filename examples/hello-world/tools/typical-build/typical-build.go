package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "hello-world",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.BuildCmdRuns{"compile"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
