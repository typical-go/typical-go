package walker

import "go/ast"

type TypeSpecListener interface {
	OnTypeSpec(*TypeSpecEvent) error
}

type TypeSpecEvent struct {
	*ast.TypeSpec
	Name     string
	Filename string
	File     *ast.File
}

func (e *TypeSpecEvent) CommentDoc() string {
	if e.Doc != nil {
		return e.Doc.Text()
	}
	return ""
}

func (e *TypeSpecEvent) Annotations() Annotations {
	return ParseAnnotations(e.CommentDoc())
}

func (e *TypeSpecEvent) IsInterface() bool {
	switch e.TypeSpec.Type.(type) {
	case *ast.InterfaceType:
		return true
	}
	return false
}
