package app

import (
	"fmt"
	"io"

	"github.com/typical-go/typical-go/examples/typmock-sample/internal/greeter"
)

// Start application
func Start(w io.Writer, g greeter.Greeter) {
	fmt.Fprintln(w, g.Greet())
}
