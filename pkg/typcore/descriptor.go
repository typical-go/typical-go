package typcore

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
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

	// ConfigManager of the project (OPTIONAL).
	ConfigManager
}

// LaunchApp to launch the app
func (d *Descriptor) LaunchApp() (err error) {
	if d.App == nil {
		return errors.New("Descriptor is missing `App`")
	}
	if err = d.Validate(); err != nil {
		return
	}
	return d.App.Run(d)
}

// LaunchBuildTool to launch the build tool
func (d *Descriptor) LaunchBuildTool() (err error) {
	if d.BuildTool == nil {
		return errors.New("Descriptor is missing `BuildTool`")
	}

	var c *Context
	if c, err = CreateContext(d); err != nil {
		return
	}

	if err = common.Validate(c); err != nil {
		return
	}

	if err = d.Precondition(&PreconditionContext{Context: c}); err != nil {
		return
	}

	return d.BuildTool.Run(c)
}

// Precondition for this project
func (d *Descriptor) Precondition(c *PreconditionContext) (err error) {
	if preconditioner, ok := c.App.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	if preconditioner, ok := c.BuildTool.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-Build-Tool: %w", err)
		}
	}

	if preconditioner, ok := c.ConfigManager.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-Config-Manager: %w", err)
		}
	}

	return
}

// Validate context
func (d *Descriptor) Validate() (err error) {
	if err = validateName(d.Name); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if d.App == nil {
		return errors.New("Descriptor: App can't be nil")
	} else if err = common.Validate(d.App); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	if d.BuildTool == nil {
		return errors.New("Descriptor: BuildTool can't be nil")
	} else if err = common.Validate(d.BuildTool); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	return
}

func validateName(name string) (err error) {
	if name == "" {
		return errors.New("Name can't be empty")
	}

	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return errors.New("Invalid name")
	}
	return
}
