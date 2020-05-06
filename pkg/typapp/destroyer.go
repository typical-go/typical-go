package typapp

import "go.uber.org/dig"

var (
	_ Destroyer = (*Destructor)(nil)
)

type (
	// Destroyer responsible to destroy the dependency [mock]
	Destroyer interface {
		Destructors() []*Destructor
	}

	// Destroyers is list destroyer
	Destroyers []Destroyer

	// Destructor is invocation to destroy dependency
	Destructor struct {
		Fn interface{}
	}
)

// Invoke the invocation using dig container
func (d *Destructor) Invoke(di *dig.Container) (err error) {
	if d.Fn == nil {
		panic("destroy: missing Fn")
	}
	return di.Invoke(d.Fn)
}

// Destructors return slice of destruction
func (d *Destructor) Destructors() []*Destructor {
	return []*Destructor{d}
}

// Destructors return slice of destruction
func (d Destroyers) Destructors() (dtors []*Destructor) {
	for _, destroyer := range d {
		dtors = append(dtors, destroyer.Destructors()...)
	}
	return
}
