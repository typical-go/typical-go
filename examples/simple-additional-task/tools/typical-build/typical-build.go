package main

import (
	"fmt"
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "simple-additional-task",
		Version: "1.0.0",
		Layouts: typgo.Layouts{"internal"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
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
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
