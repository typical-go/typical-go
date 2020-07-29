package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{},
		},

		// run
		&typgo.RunCmd{
			Before: typgo.BuildSysRuns{"compile"},
			Action: &typgo.StdRun{},
		},

		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
