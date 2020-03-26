package helloworld

import (
	"fmt"
	"io"
	"os"
)

// Main function to run hello-world
func Main(greeter Greeter, w io.Writer) {
	fmt.Fprintln(w, greeter.Greet())
}

// GetWriter to get writer to greet the world [constructor]
func GetWriter() io.Writer {
	return os.Stdout
}
