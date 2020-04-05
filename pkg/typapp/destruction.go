package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

var (
	_ Destroyer = (*Destruction)(nil)
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

// Destructions return slice of destruction
func (d *Destruction) Destructions() []*Destruction {
	return []*Destruction{d}
}
