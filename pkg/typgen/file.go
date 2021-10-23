package typgen

import (
	"strings"
)

type (
	// File information
	File struct {
		Path    string
		Name    string
		Imports []*Import
	}
)

var _ Coder = (*File)(nil)

func (f *File) Code() string {
	var b strings.Builder
	b.WriteString("package ")
	b.WriteString(f.Name)
	b.WriteString("\n")

	b.WriteString("\nimport (\n")
	for _, i := range f.Imports {
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
