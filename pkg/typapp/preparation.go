package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

var (
	_ Preparer = (*Preparation)(nil)
)

// Preparation is invocation to prepare the application
type Preparation struct {
	*typdep.Invocation
}

// NewPreparation return new isntance of Preparation
func NewPreparation(fn interface{}) *Preparation {
	return &Preparation{
		Invocation: typdep.NewInvocation(fn),
	}
}

// Preparations return preparation as its slice
func (p *Preparation) Preparations() []*Preparation {
	return []*Preparation{p}
}
