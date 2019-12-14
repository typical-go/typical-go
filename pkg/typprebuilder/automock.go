package typprebuilder

import (
	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

type Automocks []string

// OnTypeSpec handle type specificatio event
func (a *Automocks) OnTypeSpec(e *walker.TypeSpecEvent) (err error) {
	if e.IsInterface() {
		annotations := e.Annotations()
		if !annotations.Contain("nomock") {
			*a = append(*a, e.Filename)
		}
	}
	return
}
