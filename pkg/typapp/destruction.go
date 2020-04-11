package typapp

import "go.uber.org/dig"

var (
	_ Destroyer = (*Destruction)(nil)
)

// Destruction is invocation to destroy dependency
type Destruction struct {
	fn interface{}
}

// NewDestruction return new instance of Destructor
func NewDestruction(fn interface{}) *Destruction {
	return &Destruction{
		fn: fn,
	}
}

// Invoke the invocation using dig container
func (d *Destruction) Invoke(di *dig.Container) (err error) {
	return di.Invoke(d.fn)
}

// Destructions return slice of destruction
func (d *Destruction) Destructions() []*Destruction {
	return []*Destruction{d}
}
