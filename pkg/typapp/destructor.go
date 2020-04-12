package typapp

import "go.uber.org/dig"

var (
	_ Destroyer = (*Destructor)(nil)
)

// Destructor is invocation to destroy dependency
type Destructor struct {
	fn interface{}
}

// NewDestructor return new instance of Destructor
func NewDestructor(fn interface{}) *Destructor {
	return &Destructor{
		fn: fn,
	}
}

// Invoke the invocation using dig container
func (d *Destructor) Invoke(di *dig.Container) (err error) {
	return di.Invoke(d.fn)
}

// Destructors return slice of destruction
func (d *Destructor) Destructors() []*Destructor {
	return []*Destructor{d}
}
