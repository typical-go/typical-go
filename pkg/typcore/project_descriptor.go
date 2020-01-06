package typcore

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
)

// ProjectDescriptor describe the project
type ProjectDescriptor struct {
	Name         string
	Description  string
	Package      string
	Version      string
	AppModule    interface{}
	Modules      common.Interfaces
	ConfigLoader ConfigLoader
	Releaser     Releaser

	MockTargets  common.Strings
	Constructors common.Interfaces
}

// Validate context
func (c *ProjectDescriptor) Validate() (err error) {
	if c.Name == "" {
		return errors.New("Context: Name can't be empty")
	}
	if c.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if c.Version == "" {
		c.Version = "0.0.1"
	}
	if c.ConfigLoader == nil {
		c.ConfigLoader = DefaultConfigLoader()
	}
	if err = c.Releaser.Validate(); err != nil {
		return fmt.Errorf("Context: Releaser: %w", err)
	}

	return
}

// AllModule return app module and modules
func (c *ProjectDescriptor) AllModule() (modules []interface{}) {
	// NOTE: modules should be before app module to make sure it readiness
	modules = append(modules, c.Modules...)
	modules = append(modules, c.AppModule)
	return
}
