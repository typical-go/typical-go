package helloworld

// Greeter responsible to greet
// @mock
type Greeter interface {
	Greet() string
}

// GreeterImpl is implementation of Greeeters
type GreeterImpl struct {
}

// NewGreeter return new insteance of Greeter [constructor]
func NewGreeter() Greeter {
	return &GreeterImpl{}
}

// Greet the world
func (g *GreeterImpl) Greet() string {
	return "Hello World"
}
