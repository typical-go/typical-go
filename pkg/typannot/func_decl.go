package typannot

import "go/ast"

type (
	// FuncDecl function declaration
	FuncDecl struct {
		Name string
		Docs []string
	}
)

var _ DeclType = (*FuncDecl)(nil)

func createFuncDecl(funcDecl *ast.FuncDecl, file File) DeclType {
	return &FuncDecl{
		Name: funcDecl.Name.Name,
		Docs: docs(funcDecl.Doc),
	}
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
