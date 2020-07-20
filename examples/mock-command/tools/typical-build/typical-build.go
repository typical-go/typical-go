package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "mock-command",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Cmds: []typgo.Cmd{

			&typgo.CompileCmd{
				Action: typgo.Actions{
					typannot.Annotators{
						&typapp.CtorAnnotation{},
					},
					&typgo.StdCompile{},
				},
			},

			&typgo.RunCmd{
				Precmds: []string{"compile"},
				Action:  &typgo.StdRun{},
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
