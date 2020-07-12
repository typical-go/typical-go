package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "generate-mock",
		Version: "1.0.0",

		Layouts: []string{"internal"},

		Compile: typgo.Compilers{
			&typgo.CtorAnnotation{},
			&typgo.StdCompile{},
		},
		Run:   &typgo.StdRun{},
		Test:  &typgo.StdTest{},
		Clean: &typgo.StdClean{},

		Utility: &typmock.Utility{},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
