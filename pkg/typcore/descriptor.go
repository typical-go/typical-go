package typcore

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
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

	constructors common.Interfaces
}

// Validate context
func (c *Descriptor) Validate() (err error) {
	if c.Version == "" {
		c.Version = "0.0.1"
	}
	if c.Name == "" {
		return errors.New("Context: Name can't be empty")
	}
	if c.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if c.Build == nil {
		return errors.New("Context: Build can't be empty")
	}
	if err = c.Build.Validate(); err != nil {
		return fmt.Errorf("Context: %w", err)
	}
	return
}
