package typprebuilder

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

// Autowires is list of function declarion for
type Autowires []string

// IsAction return true if function declaration is eligble for autowrire
func (a *Autowires) IsAction(e *walker.FuncDeclEvent) bool {
	var godoc string
	if e.Doc != nil {
		godoc = e.Doc.Text()
	}
	annotations := walker.ParseAnnotations(godoc)
	if strings.HasPrefix(e.Name, "New") {
		return !annotations.Contain("nowire")
	}
	return annotations.Contain("autowire")

}

// ActionPerformed is when function is autowired
func (a *Autowires) ActionPerformed(e *walker.FuncDeclEvent) (err error) {
	*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.Name))
	return
}
