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
			Usage: `print "pong"`,
			Action: typgo.NewAction(func(c *typgo.Context) error {
				// new action with golang implementation
				fmt.Println("pong")
				return nil
			}),
		},
		// gov
		&typgo.Task{
			Name:  "gov",
			Usage: "print go version",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				// you can also call bash command
				return c.Execute(&typgo.Bash{
					Name:   "go",
					Args:   []string{"version"},
					Stdout: os.Stdout,
					Stderr: os.Stderr,
					Stdin:  os.Stdin,
				})
			}),
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
