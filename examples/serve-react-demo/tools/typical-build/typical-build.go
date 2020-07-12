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
		Layouts: []string{"internal"},

		Compile: typgo.Compilers{
			typgo.NewCompiler(npmBuild),
			&typgo.StdCompile{},
		},
		Run:   &typgo.StdRun{},
		Clean: &typgo.StdClean{},
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
