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

	App           AppInterface
	Build         BuildInterface
	Configuration ConfigurationInterface

	constructors common.Interfaces
}

// Validate context
func (c *Descriptor) Validate() (err error) {
	if c.Name == "" {
		return errors.New("Context: Name can't be empty")
	}
	if c.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if c.Version == "" {
		c.Version = "0.0.1"
	}
	if c.Build != nil {
		if err = c.Build.Validate(); err != nil {
			return fmt.Errorf("Context: %w", err)
		}
	}
	return
}

// AppendConstructor to append constructor
func (c *Descriptor) AppendConstructor(constructors ...interface{}) {
	c.constructors.Append(constructors...)
}

// Constructors return contruction functions
func (c *Descriptor) Constructors() []interface{} {
	return c.constructors.Slice()
}