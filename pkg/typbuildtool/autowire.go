package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Autowires is list of function declarion to be provided by dig
type Autowires []string

// OnFuncDecl is when function is autowired
func (a *Autowires) OnFuncDecl(e *walker.FuncDeclEvent) (err error) {
	annotations := e.Annotations()
	if annotations.Contain("autowire") {
		*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.Name))
	}
	return
}
