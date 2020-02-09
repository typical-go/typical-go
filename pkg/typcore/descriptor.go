package typcore

import (
	"errors"
	"fmt"

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
		return errors.New("Context: Name can't be empty")
	}
	if d.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if d.Build == nil {
		return errors.New("Context: Build can't be empty")
	}

	//
	// Validate object field
	//
	if err = d.Build.Validate(); err != nil {
		return fmt.Errorf("Context: %w", err)
	}
	return
}

// BuildContext return build context of descriptor
func (d *Descriptor) BuildContext() (bctx *BuildContext, err error) {
	var (
		projInfo     *ProjectInfo
		declarations []*walker.Declaration
	)
	if err = d.Validate(); err != nil {
		return
	}
	if projInfo, err = ReadProject(typenv.Layout.App); err != nil {
		return
	}
	if declarations, err = walker.Walk(projInfo.Files); err != nil {
		return
	}
	return &BuildContext{
		Descriptor:   d,
		Declarations: declarations,
		ProjectInfo:  projInfo,
	}, nil
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
