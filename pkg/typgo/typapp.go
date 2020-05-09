package typgo

var (
	_ctors []*Constructor
	_dtors []*Destructor
)

// Provide constructor
func Provide(ctors ...*Constructor) {
	_ctors = append(_ctors, ctors...)
}

// Destroy destructor
func Destroy(dtors ...*Destructor) {
	_dtors = append(_dtors, dtors...)
}
