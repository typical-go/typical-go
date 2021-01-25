package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-task",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.TaskNames{"build"},
		},
		// ping
		&typgo.Task{
			Name:  "ping",
			Usage: `print "pong"`,
			Action: typgo.NewAction(func(c *typgo.Context) error {
				fmt.Println("pong") // new action with golang implementation
				return nil
			}),
		},
		// gov
		&typgo.Task{
			Name:  "gov",
			Usage: "print go version",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				return c.ExecuteBash("go version") // you can also call bash command
			}),
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
