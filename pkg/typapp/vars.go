package typapp

import (
	"io"
	"os"
)

var (
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)
