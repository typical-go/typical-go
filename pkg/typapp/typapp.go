package typapp

var (
	_ctors []*Constructor
)

// Provide constructor globally
func Provide(cons ...*Constructor) {
	_ctors = append(_ctors, cons...)
}
