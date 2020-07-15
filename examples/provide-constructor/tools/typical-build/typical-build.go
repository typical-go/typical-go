package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "provide-constructor",
		Version: "1.0.0",
		Layouts: typgo.Layouts{"internal"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: &typgo.Actions{
					&typgo.CtorAnnotation{},
					&typgo.DtorAnnotation{},
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

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
