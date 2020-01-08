package walker

import "go/ast"

// DeclType is declaration type
type DeclType int

const (
	FuncDeclType DeclType = iota + 1
	InterfaceSpecType
	GenSpecType
)

// DeclListener listen declaration event
type DeclListener interface {
	OnDecl(*DeclEvent) error
}

// DeclEvent happen when declarion
type DeclEvent struct {
	Name     string
	Filename string
	File     *ast.File
	Doc      Doc
	Type     DeclType
	Source   interface{}
}

// Doc is go documentation
type Doc string

// Annotations of doc
func (d Doc) Annotations() (annotations Annotations) {
	if d == "" {
		return
	}
	return ParseAnnotations(string(d))
}
