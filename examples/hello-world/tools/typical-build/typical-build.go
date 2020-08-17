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
		// clean
		&typgo.CleanProject{},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"compile"},
			Action: &typgo.RunProject{},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
