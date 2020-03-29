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

// Prepare the dependency
func (p *Preparation) Prepare() []*Preparation {
	return []*Preparation{p}
}
