package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "custom-command",
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

			&typgo.Command{
				Name: "ping",
				Action: typgo.NewAction(func(c *typgo.Context) error {
					fmt.Println("pong")
					return nil
				}),
			},
		},
	}
)

func main() {
	typgo.Start(&descriptor)
}
