package app

import (
	"fmt"

	"go.uber.org/dig"
)

type (
	// @ctor
	SomeStruct struct{}
	// @ctor
	SomeInterface interface{}
)

// Start the application which invoked from main() function in cmd package.
func Start(di *dig.Container, text string) {
	// "text" is provided by dependency-injection
	fmt.Println(text)

	// Learn more: https://godoc.org/go.uber.org/dig#hdr-Named_Values
	type parameter struct {
		dig.In
		Text string `name:"typical"`
	}

	// Invoke another function
	di.Invoke(func(p parameter) {
		fmt.Println(p.Text)
	})
}

// Stop the application which invoked gracefull when the application stop or received exit signal
func Stop() { fmt.Println("Bye") }

// HelloWorld text
// @ctor
func HelloWorld() string { return "Hello World" }

// HelloTypical text
// @ctor (name:"typical")
func HelloTypical() string { return "Hello Typical" }
