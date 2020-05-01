package typapp

var (
	global []*Constructor
)

// Provide constructor globally
func Provide(cons ...*Constructor) {
	global = append(global, cons...)
}
