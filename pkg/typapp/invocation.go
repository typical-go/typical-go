package typapp

import "go.uber.org/dig"

// Invocation detail
type Invocation struct {
	fn interface{}
}

// NewInvocation return new instance of invocation
func NewInvocation(fn interface{}) *Invocation {
	return &Invocation{
		fn: fn,
	}
}

// Invoke the invocation using dig container
func (i *Invocation) Invoke(di *dig.Container) (err error) {
	return di.Invoke(i.fn)
}
