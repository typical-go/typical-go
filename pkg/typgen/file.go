package typgen

import (
	"go/ast"
	"path/filepath"
	"strings"
)

type (
	// File information
	File struct {
		Path   string
		Name   string
		Import []*Import
	}
	Import struct {
		Name string
		Path string
	}
)

var _ Coder = (*File)(nil)

func CreateImports(f *ast.File) []*Import {
	var imports []*Import
	for _, i := range f.Imports {
		imports = append(imports, createImport(i))
	}
	return imports
}

func PackageName(path string) string {
	return filepath.Base(filepath.Dir(path))
}

func createImport(i *ast.ImportSpec) *Import {
	name := ""
	if i.Name != nil {
		name = i.Name.Name
	}
	return &Import{
		Name: name,
		Path: i.Path.Value,
	}
}

func (f *File) Code() string {
	var b strings.Builder
	b.WriteString("package ")
	b.WriteString(f.Name)
	b.WriteString("\n")

	b.WriteString("\nimport (\n")
	for _, i := range f.Import {
		b.WriteString("\t")
		if i.Name != "" {
			b.WriteString(i.Name)
			b.WriteString(" ")
		}
		b.WriteString("\"")
		b.WriteString(i.Path)
		b.WriteString("\"")
		b.WriteString("\n")
	}
	b.WriteString(")")
	return b.String()
}
