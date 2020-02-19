package typcore

import (
	"errors"
	"path/filepath"
)

// Descriptor describe the project
type Descriptor struct {
	Name        string
	Description string
	Package     string
	Version     string
	Sources     []string

	App           App
	BuildTool     BuildTool
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
