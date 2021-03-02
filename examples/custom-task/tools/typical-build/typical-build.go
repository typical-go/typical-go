package main

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-task",
	ProjectVersion: "1.0.0",

	Environment: typgo.Environment{
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
				c.Info("pong") // new action with golang implementation
				return nil
			}),
		},
		// info
		&typgo.Task{
			Name:  "info",
			Usage: "print info",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				c.Info("print the info:")
				c.ExecuteBash("go version")
				c.ExecuteBash("git version")
				c.Infof("\nENV: key1=%s\n", os.Getenv("key1"))
				return nil
			}),
		},
		// help
		&typgo.Task{
			Name:  "go-help",
			Usage: "print go help",
			Action: &typgo.Bash{
				Name:   "go",
				Args:   []string{"help"},
				Stdout: os.Stdout,
			},
		},
		// multi-task
		&typgo.Task{
			Name:   "multi-task",
			Usage:  "run multi-task",
			Action: typgo.TaskNames{"ping", "info"},
		},
		// database
		&typgo.Task{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database tool",
			SubTasks: []*typgo.Task{
				{
					Name:  "create",
					Usage: "create database",
					Action: typgo.NewAction(func(c *typgo.Context) error {
						c.Info("create database")
						return nil
					}),
				},
				{
					Name:  "drop",
					Usage: "drop database",
					Action: typgo.NewAction(func(c *typgo.Context) error {
						c.Info("drop databse")
						return nil
					}),
				},
				{
					Name:  "migrate",
					Usage: "migrate database",
					Action: typgo.NewAction(func(c *typgo.Context) error {
						c.Info("migrate databse")
						return nil
					}),
				},
				{
					Name:  "seed",
					Usage: "seed database",
					Action: typgo.NewAction(func(c *typgo.Context) error {
						c.Info("seed databse")
						return nil
					}),
				},
			},
		},
		// greet
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
	c.Infof("Hello %s\n", g.person)
	return nil
}

func main() {
	typgo.Start(&descriptor)
}
