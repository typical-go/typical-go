package greeter

type (
	// Greeter responsible to greet
	// @mock
	Greeter interface {
		Greet() string
	}
	greeterImpl struct{}
)

// New insteance of Greeter
func New() Greeter {
	return &greeterImpl{}
}

// Greet the world
func (g *greeterImpl) Greet() string {
	return "Hello World"
}
