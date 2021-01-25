package main

import (
	"fmt"
	"os"

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
			Usage: "print pong",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				fmt.Println("pong") // new action with golang implementation
				return nil
			}),
		},
		// info
		&typgo.Task{
			Name:  "info",
			Usage: "print info",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				fmt.Println("print the info:")
				c.ExecuteBash("go version")
				c.ExecuteBash("git version")
				return nil
			}),
		},
		// help
		&typgo.Task{
			Name:  "help",
			Usage: "print help",
			Action: &typgo.Bash{
				Name:   "go",
				Args:   []string{"help"},
				Stdout: os.Stdout,
			},
		},
		&typgo.Task{
			Name:   "all",
			Usage:  "run all custom task",
			Action: typgo.TaskNames{"ping", "info", "help"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
