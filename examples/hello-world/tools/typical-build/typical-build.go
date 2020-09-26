package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "hello-world",
	ProjectVersion: "1.0.0",

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileProject{},
		// run
		&typgo.RunProject{
			Before: typgo.BuildCmdRuns{"compile"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
