package typapp

import (
	"io"
	"os"

	"go.uber.org/dig"
)

type (
	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

var (
	glob []*Constructor
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)

// Provide constructor
func Provide(name string, fn interface{}) {
	glob = append(glob, &Constructor{Name: name, Fn: fn})
}

// Container for dependency injection
func Container() (*dig.Container, error) {
	di := dig.New()
	for _, c := range glob {
		if err := di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return nil, err
		}
	}
	return di, nil
}

// Reset constructor
func Reset() {
	glob = make([]*Constructor, 0)
}
