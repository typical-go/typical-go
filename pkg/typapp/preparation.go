package typapp

import "go.uber.org/dig"

var (
	_ Preparer = (*Preparation)(nil)
)

// Preparation is invocation to prepare the application
type Preparation struct {
	fn interface{}
}

// NewPreparation return new isntance of Preparation
func NewPreparation(fn interface{}) *Preparation {
	return &Preparation{
		fn: fn,
	}
}

// Preparations return preparation as its slice
func (p *Preparation) Preparations() []*Preparation {
	return []*Preparation{p}
}

// Invoke the invocation using dig container
func (p *Preparation) Invoke(di *dig.Container) (err error) {
	return di.Invoke(p.fn)
}
