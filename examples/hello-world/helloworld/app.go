package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

// Main function of hello-world
func Main(d *typgo.Descriptor) (err error) {
	fmt.Println("Hello World")
	return
}
