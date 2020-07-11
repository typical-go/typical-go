package main

import (
	"log"

	"github.com/typical-go/typical-go/examples/configuration-with-invocation/internal/server"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "configuration-with-invocation",
		Version: "1.0.0",

		Layouts: []string{"internal"},

		Compile: &typgo.StdCompile{
			Before: &typgo.ConfigManager{
				Configs: []*typgo.Configuration{
					{Name: "SERVER", Spec: &server.Config{}},
				},
			},
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
