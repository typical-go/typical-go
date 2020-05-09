package typapp

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
