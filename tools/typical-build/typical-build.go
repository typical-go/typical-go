package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "typical-go",
		Version: "0.9.57",
		Layouts: typgo.Layouts{"internal", "pkg"},

		Cmds: typgo.Cmds{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{MainPackage: "."},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
			},
			&typgo.TestCmd{
				Action: &typgo.StdTest{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},

			&typgo.Command{
				Name:    "examples",
				Aliases: []string{"e"},
				Usage:   "Test all example",
				Action:  typgo.NewAction(testExamples),
			},

			&typrls.Command{
				Validation: typrls.DefaultValidation,
				Summary:    typrls.DefaultSummary,
				Releaser:   &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
			},
		},
	}
)

func testExamples(c *typgo.Context) error {
	return c.Execute(&execkit.GoTest{
		Packages: []string{"./examples/..."},
	})
}

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err.Error())
	}
}
