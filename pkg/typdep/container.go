package typdep

import (
	"go.uber.org/dig"
)

// Container of dependency
type Container struct {
	container *dig.Container
}

// New container
func New() *Container {
	return &Container{
		container: dig.New(),
	}
}
