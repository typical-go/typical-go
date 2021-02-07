package main

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-task",
	ProjectVersion: "1.0.0",

	EnvLoader: typgo.EnvMap{
		"key1": "value1",
	},

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
				fmt.Printf("\nENV: key1=%s\n", os.Getenv("key1"))
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
			Action: typgo.TaskNames{"ping", "info"},
		},
		&greetTask{person: "john doe"},
	},
}

type greetTask struct {
	person string
}

var _ typgo.Tasker = (*greetTask)(nil)
var _ typgo.Action = (*greetTask)(nil)

func (g *greetTask) Task() *typgo.Task {
	return &typgo.Task{
		Name:   "greet",
		Usage:  "greet person",
		Action: g,
	}
}

func (g *greetTask) Execute(c *typgo.Context) error {
	fmt.Println("Hello " + g.person)
	return nil
}

func main() {
	typgo.Start(&descriptor)
}
