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

// AppendCtor append contructor to global variable
func AppendCtor(ctors ...*Constructor) {
	_ctors = append(_ctors, ctors...)
}

// AppendDtor append destructor to global variable
func AppendDtor(dtors ...*Destructor) {
	_dtors = append(_dtors, dtors...)
}

// GetCtors return list of global constructor
func GetCtors() []*Constructor {
	return _ctors
}

// GetDtors return list of global destructor
func GetDtors() []*Destructor {
	return _dtors
}
