package main

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.11.5",

	Tasks: []typgo.Tasker{
		// compile
		&typgo.GoBuild{MainPackage: "."},
		// run
		&typgo.RunBinary{
			Before: typgo.TaskNames{"build"},
		},
		// test
		&typgo.GoTest{
			Includes: []string{"internal/**", "pkg/**"},
		},
		// test-examples
		&typgo.Task{
			Name:    "test-examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action: &typgo.Command{
				Name:   "go",
				Args:   []string{"test", "./examples/..."},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			},
		},
		// test-setup
		&typgo.Task{
			Name:  "test-setup",
			Usage: "test setup command",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				os.RemoveAll("examples/my-project")
				err := c.ExecuteCommand(&typgo.Command{
					Name:   "../bin/typical-go",
					Args:   []string{"setup", "-new", "-go-mod", "-project-pkg=github.com/typical-go/typical-go/examples/my-project"},
					Dir:    "examples",
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				})
				if err != nil {
					return err
				}

				os.RemoveAll("examples/my-project/go.mod")
				os.RemoveAll("examples/my-project/go.sum")
				return c.ExecuteCommand(&typgo.Command{
					Name:   "./typicalw",
					Args:   []string{"run"},
					Dir:    "examples/my-project",
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				})
			}),
		},
		// test-setup
		&typgo.Task{
			Name:   "test-all",
			Usage:  "test project, test examples and  test setup command",
			Action: typgo.TaskNames{"test", "build", "test-examples", "test-setup"},
		},
		// release
		&typrls.ReleaseProject{
			Before:    typgo.TaskNames{"test-all"},
			Publisher: &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
