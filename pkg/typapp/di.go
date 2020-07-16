package typapp

type (
	// Destructor is invocation to destroy dependency
	Destructor struct {
		Fn interface{}
	}
	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

var (
	_ctors []*Constructor
	_dtors []*Destructor
)

// AppendConstructor append contructor to global variable
func AppendConstructor(ctors ...*Constructor) {
	_ctors = append(_ctors, ctors...)
}

// AppendDestructor append destructor to global variable
func AppendDestructor(dtors ...*Destructor) {
	_dtors = append(_dtors, dtors...)
}

// GetConstructors return list of global constructor
func GetConstructors() []*Constructor {
	return _ctors
}

// GetDestructors return list of global destructor
func GetDestructors() []*Destructor {
	return _dtors
}
