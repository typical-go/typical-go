package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Main function of hello-world
func Main(d *typcore.Descriptor) error {
	fmt.Println("Hello World")
	return nil
}
