package typast

import (
	"go/ast"
)

type (
	// FuncDecl function declaration
	FuncDecl struct {
		Name   string
		Docs   []string
		Recv   *FieldList
		Params *FieldList
	}
	// FieldList function parameter
	FieldList struct {
		List []*Field
	}
)

var _ Type = (*FuncDecl)(nil)

func createFuncDecl(funcDecl *ast.FuncDecl, file File) Type {
	var recv, params *FieldList
	if funcDecl.Recv != nil {
		recv = createFuncParam(funcDecl.Recv)
	}
	if funcDecl.Type.Params != nil {
		params = createFuncParam(funcDecl.Type.Params)
	}
	return &FuncDecl{
		Name:   funcDecl.Name.Name,
		Docs:   docs(funcDecl.Doc),
		Recv:   recv,
		Params: params,
	}
}

func createFuncParam(l *ast.FieldList) *FieldList {
	var list []*Field
	if l != nil {
		for _, f := range l.List {
			list = append(list, createField(f))
		}
	}
	return &FieldList{List: list}
}

//
// FuncDecl
//

// GetName of declaration
func (f *FuncDecl) GetName() string {
	return f.Name
}

// GetDocs comment documentation
func (f *FuncDecl) GetDocs() []string {
	return f.Docs
}
