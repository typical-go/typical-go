package typapp

import "go.uber.org/dig"

type (
	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

var (
	constructors []*Constructor
	container    *dig.Container
)

// Provide constructor
func Provide(name string, fn interface{}) {
	constructors = append(constructors, &Constructor{Name: name, Fn: fn})
}

func Reset() {
	SetConstructors(nil)
	SetContainer(nil)
}

func SetConstructors(c []*Constructor) {
	constructors = c
}

func SetContainer(c *dig.Container) {
	container = c
}

func Constructors() []*Constructor {
	return constructors
}

func Container() (*dig.Container, error) {
	if container == nil {
		container = dig.New()
		container.Provide(func() *dig.Container { return container })
		for _, c := range constructors {
			if err := container.Provide(c.Fn, dig.Name(c.Name)); err != nil {
				return nil, err
			}
		}
	}
	return container, nil
}

func Invoke(fn interface{}) error {
	c, err := Container()
	if err != nil {
		return err
	}
	return c.Invoke(fn)
}
