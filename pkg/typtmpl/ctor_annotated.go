package typtmpl

import (
	"fmt"
	"io"

	"github.com/typical-go/typical-go/pkg/typast"
)

type (
	// CtorAnnotated template
	CtorAnnotated struct {
		Package string
		Imports []string
		Ctors   []*Ctor
	}
	// Ctor is constructor model
	Ctor struct {
		Name string `json:"name"`
		Def  string `json:"-"`
	}
)

//
// Ctor
//

// CreateCtor to create new instance of Ctor
func CreateCtor(annot *typast.Annotation) (*Ctor, error) {
	var ctor Ctor
	if err := annot.Unmarshal(&ctor); err != nil {
		return nil, fmt.Errorf("%s: %w", annot.Decl.Name, err)
	}
	ctor.Def = fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name)
	return &ctor, nil
}

//
// CtorAnnotated
//

const ctorGenerated = `package {{.Package}}

// Autogenerated by Typical-Go. DO NOT EDIT.

import ({{range $import := .Imports}}
	"{{$import}}"{{end}}
)

func init() { {{if .Ctors}}
	typapp.AppendCtor({{range $c := .Ctors}}
		&typapp.Constructor{Name: "{{$c.Name}}", Fn: {{$c.Def}}},{{end}}
	){{end}}
}`

// Execute app precondition template
func (t *CtorAnnotated) Execute(w io.Writer) (err error) {
	return Parse("ctorGenerated", ctorGenerated, t, w)
}