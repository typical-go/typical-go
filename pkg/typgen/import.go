package typgen

import (
	"go/ast"
	"strings"
)

type (
	Import struct {
		Name string
		Path string
	}
)

func CreateImports(f *ast.File) []*Import {
	var imports []*Import
	for _, i := range f.Imports {
		imports = append(imports, createImport(i))
	}
	return imports
}

func createImport(i *ast.ImportSpec) *Import {
	name := ""
	if i.Name != nil {
		name = i.Name.Name
	}
	path := strings.Trim(i.Path.Value, "\"")
	return &Import{
		Name: name,
		Path: path,
	}
}
