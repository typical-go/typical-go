package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

var (
	global []*Constructor
)

var (
	_ Provider = (*Constructor)(nil)
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

// Constructors is list of constructor
func (c *Constructor) Constructors() []*Constructor {
	return []*Constructor{c}
}

// AppendConstructor to append constructor
func AppendConstructor(cons ...*Constructor) {
	global = append(global, cons...)
}
