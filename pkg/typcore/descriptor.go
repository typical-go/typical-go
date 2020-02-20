package typcore

import (
	"errors"
	"path/filepath"
)

// Descriptor describe the project
type Descriptor struct {

	// Name of the project (OPTIONAL)
	// It should be a characters with/without underscore or dash.
	// By default, project name is same with project folder
	Name string

	// Description of the project (OPTIONAL)
	Description string

	// Version of the project (MANDATORY)
	Version string

	// Package of the project (MANDATORY)
	// It should be same with go.mod file
	Package string

	// Sources of project (OPTIONAL)
	// If `App` is sourceable, the sources will be take from it.
	// Else, by default, sources will be package name of `App` type
	// If `pkg` folder available, it will be added to sources.
	Sources []string

	// App of the project (MANDATORY)
	App App

	// BuildTool of the project (MANDATORY)
	BuildTool BuildTool

	// Configuration of the project (OPTIONAL)
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
	c := &TypicalContext{
		Descriptor:    d,
		ProjectLayout: DefaultLayout,
		Dirs:          d.Sources,
	}
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
