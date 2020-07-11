package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "hello-world",
		Version: "1.0.0",

		Compile: &typgo.StdCompile{},
		Run:     &typgo.StdRun{},
		Clean:   &typgo.StdClean{},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
