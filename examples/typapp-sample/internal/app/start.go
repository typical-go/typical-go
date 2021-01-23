package app

import (
	"fmt"

	"go.uber.org/dig"
)

type (
	// NOTE: this example of parameter object to get to named dependency
	// Learn more: https://godoc.org/go.uber.org/dig#hdr-Named_Values
	parameter struct {
		dig.In
		Text string `name:"typical"`
	}
)

// Start the application which invoked from main() function in cmd package.
func Start(di *dig.Container, text string) {
	// NOTE: We are using dependency injection powered by dig library. The function parameter is initiating by dig container.
	// Learn more: https://godoc.org/go.uber.org/dig
	fmt.Println(text)

	// NOTE: the dig container is also provided and can be use to invoke another function
	di.Invoke(func(p parameter) {
		fmt.Println(p.Text)
	})
}

// Shutdown the application which invoked gracefull when the application stop or received exit signal
func Shutdown() { fmt.Println("Bye") }

// HelloWorld text
// @ctor
// NOTE: To provide new constructor by simply to `@ctor` in function comment
func HelloWorld() string { return "Hello World" }

// HelloTypical text
// @ctor (name:"typical")
func HelloTypical() string { return "Hello Typical" }
