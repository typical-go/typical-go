package typgo

type (
	// Prebuilder return new isntance of prebuilder
	Prebuilder interface {
		Prebuild(*Context) error
	}

	// Prebuilds composite prebuild
	Prebuilds []Prebuilder

	// PrebuildFn function
	PrebuildFn func(*Context) error

	prebuilderImpl struct {
		fn PrebuildFn
	}
)

//
// prebuilderImpl
//

// NewPrebuild return new instance of Prebuilder
func NewPrebuild(fn PrebuildFn) Prebuilder {
	return &prebuilderImpl{fn: fn}
}

func (p *prebuilderImpl) Prebuild(c *Context) error {
	return p.fn(c)
}

//
// Prebuilds
//

// Prebuild prebuilds
func (p Prebuilds) Prebuild(c *Context) error {
	for _, prebuild := range p {
		if err := prebuild.Prebuild(c); err != nil {
			return err
		}
	}
	return nil
}
