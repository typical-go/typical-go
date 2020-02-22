package typcore

import (
	"errors"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Descriptor describe the project
type Descriptor struct {

	// Name of the project (OPTIONAL).
	// It should be a characters with/without underscore or dash.
	// By default, project name is same with project folder
	Name string

	// Description of the project (OPTIONAL).
	Description string

	// Version of the project (OPTIONAL).
	// By default it is 0.0.1
	Version string

	// App of the project (MANDATORY).
	App App

	// BuildTool of the project (MANDATORY).
	BuildTool BuildTool

	// Configuration of the project (OPTIONAL).
	Configuration Configuration
}

// RunApp to run app
func (d *Descriptor) RunApp() (err error) {
	if d.App == nil {
		return errors.New("Descriptor is missing `App`")
	}
	if err = d.Validate(); err != nil {
		return
	}
	return d.App.Run(d)
}

// RunBuild to run build
func (d *Descriptor) RunBuild() (err error) {
	if d.BuildTool == nil {
		return errors.New("Descriptor is missing `BuildTool`")
	}
	if err = d.Validate(); err != nil {
		return
	}
	projectSources := DefaultProjectSources(d)
	c := &TypicalContext{
		Descriptor:     d,
		ProjectLayout:  DefaultLayout,
		Dirs:           projectSources,
		ModulePackage:  DefaultModulePackage,
		ProjectSources: projectSources,
	}
	log.Infof("Scan the project directories: %v", projectSources)
	for _, dir := range c.Dirs {
		if err = filepath.Walk(dir, c.addFile); err != nil {
			return
		}
	}
	if d.Configuration != nil {
		if err = d.Configuration.Setup(); err != nil {
			return
		}
	}
	return d.BuildTool.Run(c)
}
