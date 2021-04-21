package typrls

import "github.com/typical-go/typical-go/pkg/typgo"

type (
	// Releaser responsible to release (put file to release folder)
	Releaser interface {
		Release(*Context) error
	}
	// Releasers for composite release
	Releasers []Releaser
	// ReleaseFn release function
	NewReleaser func(*Context) error

	// Context contain data for release
	Context struct {
		*typgo.Context
		Alpha         bool
		TagName       string
		ReleaseFolder string
		Summary       string
	}
)

//
// NewReleaser
//

var _ Releaser = (NewReleaser)(nil)

func (r NewReleaser) Release(c *Context) error {
	return r(c)
}

//
// Releaser
//

var _ Releaser = (Releasers)(nil)

// Release the releasers
func (r Releasers) Release(c *Context) (err error) {
	for _, releaser := range r {
		if err = releaser.Release(c); err != nil {
			return
		}
	}
	return
}
