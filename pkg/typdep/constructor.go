package typdep

// Constructor details
type Constructor struct {
	fn interface{}
}

// NewConstructor return new instance of constructor
func NewConstructor(fn interface{}) *Constructor {
	return &Constructor{
		fn: fn,
	}
}

// Provide the constructor to dig container
func (c *Constructor) Provide(di *Container) (err error) {
	return di.container.Provide(c.fn)
}

// Fn is function of constructor
func (c *Constructor) Fn() interface{} {
	return c.fn
}
