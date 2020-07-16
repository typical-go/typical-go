package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "generate-mock",
		Version: "1.0.0",
		Layouts: typgo.Layouts{"internal"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: typgo.Actions{
					&typapp.CtorAnnotation{},
					&typgo.StdCompile{},
				},
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
			&typmock.Command{},
		},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
