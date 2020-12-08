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
			Name: "ping",
			Action: typgo.NewAction(
				func(c *typgo.Context) error {
					fmt.Println("pong")
					return nil
				},
			),
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
