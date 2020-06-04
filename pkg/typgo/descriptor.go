package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"go.uber.org/dig"
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

		EntryPoint interface{}
		Layouts    []string

		Prebuild Prebuilder
		Test     Tester
		Compile  Compiler
		Run      Runner
		Release  Releaser
		Utility  Utility
	}
)

var _ typcore.AppLauncher = (*Descriptor)(nil)
var _ typcore.BuildLauncher = (*Descriptor)(nil)

// LaunchApp to launch the app
func (d *Descriptor) LaunchApp() (err error) {
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		_, err = LoadConfig(configFile)
	}

	di := dig.New()
	if err = setDependencies(di, d); err != nil {
		return
	}

	errs := common.GracefulRun(start(di, d), stop(di))
	return errs.Unwrap()
}

// LaunchBuild to launch the build tool
func (d *Descriptor) LaunchBuild() (err error) {
	return launchBuild(d)
}
