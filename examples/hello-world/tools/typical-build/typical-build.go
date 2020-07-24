package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "hello-world",
		Version: "1.0.0",

		Cmds: []typgo.Cmd{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{},
			},
			&typgo.RunCmd{
				Precmds: []string{"compile"},
				Action:  &typgo.StdRun{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
		},
	}
)

func main() {
	typgo.Start(&descriptor)
}
