package walker

import (
	"go/ast"
)

// TypeSpecListener handle when type specification
type TypeSpecListener interface {
	OnTypeSpec(*TypeSpecEvent) error
}

// TypeSpecEvent is type specification event
type TypeSpecEvent struct {
	*ast.TypeSpec
	Name     string
	Filename string
	File     *ast.File
	GenDecl  *ast.GenDecl
}

// CommentDoc return comment document of type
func (e *TypeSpecEvent) CommentDoc() string {
	if e.GenDecl.Doc != nil {
		return e.GenDecl.Doc.Text()
	}
	return ""
}

// Annotations of type
func (e *TypeSpecEvent) Annotations() Annotations {
	doc := e.CommentDoc()
	return ParseAnnotations(doc)
}

// IsInterface return true if type is interface type
func (e *TypeSpecEvent) IsInterface() bool {
	switch e.TypeSpec.Type.(type) {
	case *ast.InterfaceType:
		return true
	}
	return false
}
