package golang

import (
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Source is source code recipe for generated.go in typical package
type Source struct {
	BuildConstraints coll.Strings
	Package          string
	Imports          coll.KeyStrings
	Structs          []Struct
	Init             *Function
	Writables        []interface {
		Write(w io.Writer) error
	}
}

// NewSource return new instance of SourceCode
func NewSource(pkg string) *Source {
	return &Source{
		Package: pkg,
		Init:    NewFunction("init"),
	}
}

func (r Source) Write(w io.Writer) (err error) {
	if len(r.BuildConstraints) > 0 {
		fmt.Fprintf(w, "// +build %s\n\n", r.BuildConstraints.Join(" "))
	}
	fmt.Fprintf(w, "package %s\n", r.Package)

	for _, ks := range r.Imports {
		fmt.Fprintln(w, ks.Format(importFormat))
	}
	for i := range r.Structs {
		if err = r.Structs[i].Write(w); err != nil {
			return
		}
	}
	if r.Init != nil && !r.Init.IsEmpty() {
		if err = r.Init.Write(w); err != nil {
			return
		}
	}
	for _, wr := range r.Writables {
		if err = wr.Write(w); err != nil {
			return
		}
	}
	return
}

// AddStruct to add struct
func (r *Source) AddStruct(structs ...Struct) *Source {
	r.Structs = append(r.Structs, structs...)
	return r
}

func importFormat(key, s string) string {
	var b strings.Builder
	b.WriteString("import ")
	if key != "" {
		b.WriteString(key)
		b.WriteString(" ")
	}
	b.WriteString("\"")
	b.WriteString(s)
	b.WriteString("\"")
	return b.String()
}
