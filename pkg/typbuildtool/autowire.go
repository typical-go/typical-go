package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Autowires is list of function declarion to be provided by dig
type Autowires []string

// OnDecl to handle declaration event
func (a *Autowires) OnDecl(e *walker.DeclEvent) (err error) {
	if e.Annotations.Contain("autowire") {
		*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.Name))
	}
	return
}
