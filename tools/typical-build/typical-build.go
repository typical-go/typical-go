package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "typical-go",
		Version: "0.9.57",

		Layouts: []string{"internal", "pkg"},

		Test: &typgo.StdTest{},
		Compile: &typgo.StdCompile{
			Source: ".",
		},
		Run:     &typgo.StdRun{},
		Clean:   &typgo.StdClean{},
		Release: &typgo.Github{Owner: "typical-go", RepoName: "typical-go"},

		Utility: typgo.CreateUtility(&cli.Command{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action: func(c *cli.Context) (err error) {
				return execkit.Run(c.Context, &execkit.GoTest{
					Targets: []string{"./examples/..."},
				})
			},
		}),
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err.Error())
	}
}
