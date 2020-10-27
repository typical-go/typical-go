package app

import (
	"fmt"

	"go.uber.org/dig"
)

type (
	greeter struct {
		dig.In
		Text string `name:"typical"`
	}
)

// Start app
func Start(text string) {
	fmt.Println(text)
}

// Start2 start app
func Start2(greeter greeter) {
	fmt.Println(greeter.Text)
}

// Shutdown app
func Shutdown() {
	fmt.Println("Bye")
}
