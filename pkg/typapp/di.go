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

// GetCtors return global constructors
func GetCtors() []*Constructor {
	return _ctors
}

// GetDtors return global destructors
func GetDtors() []*Destructor {
	return _dtors
}

// ClearCtors clear global constructors
func ClearCtors() {
	_ctors = make([]*Constructor, 0)
}

// ClearDtors clear global destructors
func ClearDtors() {
	_dtors = make([]*Destructor, 0)
}
