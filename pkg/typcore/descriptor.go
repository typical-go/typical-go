package typcore

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// Descriptor describe the project
type Descriptor struct {
	Name        string
	Description string
	Package     string
	Version     string

	App           App
	Build         Build
	Configuration Configuration
}

// Validate context
func (d *Descriptor) Validate() (err error) {
	//
	// Default value
	//
	if d.Version == "" {
		d.Version = "0.0.1"
	}

	//
	// Mandatory field
	//
	if d.Name == "" {
		return errors.New("Descriptor: Name can't be empty")
	}
	if d.Package == "" {
		return errors.New("Descriptor: Package can't be empty")
	}

	//
	// Validate object field
	//
	if err = Validate(d.Build); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}
	if err = Validate(d.App); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}
	return
}

// BuildContext return build context of descriptor
func (d *Descriptor) BuildContext() (c *BuildContext, err error) {
	root := typenv.Layout.App

	c = &BuildContext{
		Descriptor: d,
		Dirs:       []string{root},
	}
	if err = d.Validate(); err != nil {
		return
	}
	if err = filepath.Walk(root, c.addFile); err != nil {
		return
	}
	if c.Declarations, err = walker.Walk(c.Files); err != nil {
		return
	}
	return c, nil
}

// AppContext return app context of descriptor
func (d *Descriptor) AppContext() (actx *AppContext, err error) {
	if err = d.Validate(); err != nil {
		return
	}
	return &AppContext{
		Descriptor: d,
	}, nil
}

// RunApp to run app
func (d *Descriptor) RunApp() (err error) {
	var (
		actx *AppContext
	)
	if d.App == nil {
		return errors.New("Descriptor is missing `App`")
	}
	if actx, err = d.AppContext(); err != nil {
		return
	}
	return d.App.Run(actx)
}

// RunBuild to run build
func (d *Descriptor) RunBuild() (err error) {
	var (
		bctx *BuildContext
	)
	if d.Build == nil {
		return errors.New("Descriptor is missing `Build`")
	}
	if bctx, err = d.BuildContext(); err != nil {
		return
	}
	if d.Configuration != nil {
		if err = d.Configuration.Setup(); err != nil {
			return
		}
	}
	return d.Build.Run(bctx)
}
