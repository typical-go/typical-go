package typast

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Context of annotation
	Context struct {
		*typgo.Context
		*Summary
		Dirs        []string
		Destination string
	}
)

// CreateImports create import line
func (c *Context) CreateImports(projPkg string, more ...string) []string {
	var imports []string
	for _, dir := range c.Dirs {
		imports = append(imports, fmt.Sprintf("%s/%s", projPkg, dir))
	}
	imports = append(imports, more...)
	return imports
}
