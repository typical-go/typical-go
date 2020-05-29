package typgo

type (
	// Releaser responsible to release
	Releaser interface {
		Release(*Context) error
	}

	// Releasers for composite release
	Releasers []Releaser

	// ReleaseFn release function
	ReleaseFn func(*Context) error

	releaserImpl struct {
		fn ReleaseFn
	}
)

var _ Releaser = (*Releasers)(nil)

//
// releaserImpl
//

// NewReleaser return new instance of Releaser
func NewReleaser(fn ReleaseFn) Releaser {
	return &releaserImpl{fn: fn}
}

func (r *releaserImpl) Release(c *Context) error {
	return r.fn(c)
}

//
// Releaser
//

// Release the releasers
func (r Releasers) Release(c *Context) (err error) {
	for _, releaser := range r {
		if err = releaser.Release(c); err != nil {
			return
		}
	}
	return
}
