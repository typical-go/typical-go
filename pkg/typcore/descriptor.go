package typcore

import (
	"errors"
	"fmt"
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
func (c *Descriptor) Validate() (err error) {
	//
	// Default value
	//
	if c.Version == "" {
		c.Version = "0.0.1"
	}

	//
	// Mandatory field
	//
	if c.Name == "" {
		return errors.New("Context: Name can't be empty")
	}
	if c.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if c.Build == nil {
		return errors.New("Context: Build can't be empty")
	}

	//
	// Validate object field
	//
	if err = c.Build.Validate(); err != nil {
		return fmt.Errorf("Context: %w", err)
	}
	return
}
