package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "provide-constructor",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Compile: typgo.Compilers{
			&typgo.CtorAnnotation{},
			&typgo.DtorAnnotation{},

			&typgo.StdCompile{},
		},

		Run:   &typgo.StdRun{},
		Clean: &typgo.StdClean{},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
