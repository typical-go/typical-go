package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Automocks is list of filename to be mocking by mockgen
type Automocks []string

// OnAnnotation handle type specificatio event
func (a *Automocks) OnAnnotation(decl *walker.Declaration, ann *walker.Annotation) (err error) {
	*a = append(*a, decl.Filename)
	return
}
