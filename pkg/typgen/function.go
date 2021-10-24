package typgen

import (
	"go/ast"
	"strings"
)

type (
	// Function declaration
	Function struct {
		Name    string
		Docs    []string
		Recv    []*Field
		Params  []*Field
		Returns []*Field
		Body    Coder
	}
)

var _ Type = (*Function)(nil)
var _ Coder = (*Function)(nil)

func CreateFuncDecl(funcDecl *ast.FuncDecl, file *File) *Function {
	var recv, params []*Field
	if funcDecl.Recv != nil {
		for _, f := range funcDecl.Recv.List {
			recv = append(recv, createField(f))
		}
	}

	if funcDecl.Type.Params != nil {
		for _, f := range funcDecl.Type.Params.List {
			params = append(params, createField(f))
		}
	}

	if funcDecl.Type.Results != nil {
		for _, f := range funcDecl.Type.Results.List {
			params = append(params, createField(f))
		}
	}

	return &Function{
		Name:   funcDecl.Name.Name,
		Docs:   docs(funcDecl.Doc),
		Recv:   recv,
		Params: params,
	}
}

// GetName of declaration
func (f *Function) GetName() string {
	return f.Name
}

// GetDocs comment documentation
func (f *Function) GetDocs() []string {
	return f.Docs
}

// IsMethod return true if function is method (has receiver argument)
func (f *Function) IsMethod() bool {
	return len(f.Recv) > 0
}

func (f *Function) Code() string {
	var o strings.Builder
	o.WriteString("func ")
	o.WriteString(f.Name)

	o.WriteString("(")
	writeFields(&o, f.Params)
	o.WriteString(")")

	if len(f.Returns) > 0 {
		o.WriteString("(")
		writeFields(&o, f.Returns)
		o.WriteString(")")
	}
	o.WriteString("{\n")

	if f.Body != nil {
		o.WriteString(f.Body.Code())
		o.WriteString("\n")
	}

	o.WriteString("}")
	return o.String()
}

func writeFields(o *strings.Builder, fields []*Field) {
	for i, field := range fields {
		if i > 0 {
			o.WriteString(",")
		}
		for j, name := range field.Names {
			if j > 0 {
				o.WriteString(",")
			}
			o.WriteString(name)
		}
		o.WriteString(" ")
		o.WriteString(field.Type)
	}
}
