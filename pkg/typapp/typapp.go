package typapp

type (
	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

var (
	glob []*Constructor
)

// Provide constructor
func Provide(name string, fn interface{}) {
	glob = append(glob, &Constructor{Name: name, Fn: fn})
}

// Reset constructor
func Reset() {
	glob = make([]*Constructor, 0)
}
