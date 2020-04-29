package helloworld

import "fmt"

// Greeter responsible to greet
type Greeter struct {
}

// NewGreeter return new insteance of Greeter
// @constructor
func NewGreeter() *Greeter {
	return &Greeter{}
}

// Greet the world
func (g *Greeter) Greet() {
	fmt.Println("Hello World")
}
