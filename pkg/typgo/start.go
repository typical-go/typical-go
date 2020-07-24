package typgo

import (
	"log"
	"os"
)

type (
	// Descriptor describe the project
	Descriptor struct {
		// Name of the project (OPTIONAL). It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		Name string
		// Version of the project (OPTIONAL). By default it is 0.0.1
		Version string
		Layouts []string
		Cmds    []Cmd
	}
)

// Start typical build-tool
func Start(d *Descriptor) {
	if err := createBuildSys(d).app().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
