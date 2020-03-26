package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

var (
	global []*Constructor
)

// Constructor to provide the dependency
type Constructor struct {
	*typdep.Constructor
}

// NewConstructor return new instance of Constructor
func NewConstructor(fn interface{}) *Constructor {
	return &Constructor{
		Constructor: typdep.NewConstructor(fn),
	}
}

// Provide the dependency
func (c *Constructor) Provide() []*Constructor {
	return []*Constructor{c}
}

// AppendConstructor to append constructor
func AppendConstructor(cons ...*Constructor) {
	global = append(global, cons...)
}
