package walker

import "go/ast"

// FuncDeclListener is listen function declarion event
type FuncDeclListener interface {
	OnFuncDecl(*FuncDeclEvent) error
}

// FuncDeclEvent is an event when function declared
type FuncDeclEvent struct {
	*ast.FuncDecl
	Name     string
	Filename string
	File     *ast.File
}

// CommentDoc return comment documentation of function
func (e *FuncDeclEvent) CommentDoc() string {
	if e.Doc != nil {
		return e.Doc.Text()
	}
	return ""
}

// Annotations of function
func (e *FuncDeclEvent) Annotations() Annotations {
	return ParseAnnotations(e.CommentDoc())
}
