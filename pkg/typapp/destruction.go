package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

// Destruction is invocation to destroy dependency
type Destruction struct {
	*typdep.Invocation
}

// NewDestruction return new instance of Destructor
func NewDestruction(fn interface{}) *Destruction {
	return &Destruction{
		Invocation: typdep.NewInvocation(fn),
	}
}

// Destroy the dependecies
func (d *Destruction) Destroy() []*Destruction {
	return []*Destruction{d}
}
