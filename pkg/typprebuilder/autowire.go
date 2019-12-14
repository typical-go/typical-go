package typprebuilder

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

// Autowires is list of function declarion to be provided by dig
type Autowires []string

// IsAutowire return true if function declaration is eligble for autowrire
func (a *Autowires) isAutowire(e *walker.FuncDeclEvent) bool {
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

// OnFuncDecl is when function is autowired
func (a *Autowires) OnFuncDecl(e *walker.FuncDeclEvent) (err error) {
	if a.isAutowire(e) {
		*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.Name))
	}
	return
}
