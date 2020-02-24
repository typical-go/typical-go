package typdep

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
func (i *Invocation) Invoke(di *Container) (err error) {
	return di.container.Invoke(i.fn)
}
