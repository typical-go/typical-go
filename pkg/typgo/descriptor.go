package typgo

import (
	"github.com/typical-go/typical-go/pkg/typcore"
)

type (

	// Descriptor describe the project
	Descriptor struct {
		// Name of the project (OPTIONAL). It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		Name string
		// Description of the project (OPTIONAL).
		Description string
		// Version of the project (OPTIONAL). By default it is 0.0.1
		Version string

		Layouts []string

		Prebuild Prebuilder
		Test     Tester
		Compile  Compiler
		Run      Runner
		Release  Releaser
		Clean    Cleaner

		Utility Utility
	}
)

var _ typcore.BuildLauncher = (*Descriptor)(nil)

// LaunchBuild to launch the build tool
func (d *Descriptor) LaunchBuild() (err error) {
	return launchBuild(d)
}
