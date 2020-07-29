package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var descriptor = typgo.Descriptor{
	Name:    "custom-command",
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

		// ping
		&typgo.Command{
			Name: "ping",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				fmt.Println("pong")
				return nil
			}),
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
