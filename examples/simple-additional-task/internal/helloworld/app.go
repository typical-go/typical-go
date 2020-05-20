package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

// Main function to run hello-world
func Main(d *typgo.Descriptor) error {
	fmt.Println("Hello World")
	return nil
}
