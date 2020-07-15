package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "server-echo-react",
		Version: "1.0.0",
		Layouts: typgo.Layouts{"internal"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: typgo.Actions{
					typgo.NewAction(npmBuild),
					&typgo.StdCompile{},
				},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
		},
	}
)

func npmBuild(c *typgo.Context) error {
	return c.Execute(&execkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  "react-demo",
	})
}

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
