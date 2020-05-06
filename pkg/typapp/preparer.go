package typapp

var (
	_ Preparer = (*Preparation)(nil)
	_ Preparer = (*Preparers)(nil)
)

type (
	// Preparer responsible to prepare the dependency[mock]
	Preparer interface {
		Preparations() []*Preparation
	}

	// Preparers return list of prepare
	Preparers []Preparer

	// Preparation is invocation to prepare the application
	Preparation struct {
		Fn interface{}
	}
)

// Preparations return preparation as its slice
func (p *Preparation) Preparations() []*Preparation {
	return []*Preparation{p}
}

// Preparations return preparation as its slice
func (p Preparers) Preparations() (preps []*Preparation) {
	for _, preparer := range p {
		preps = append(preps, preparer.Preparations()...)
	}
	return
}
