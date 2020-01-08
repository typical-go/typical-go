package walker

import "go/ast"

// EventType is declaration type
type EventType string

const (
	FunctionType  = EventType("Function")
	InterfaceType = EventType("Interface")
	StructType    = EventType("Struct")
	GenericType   = EventType("Generic")
)

// DeclListener listen declaration event
type DeclListener interface {
	OnDecl(*DeclEvent) error
}

// DeclEvent happen when declarion
type DeclEvent struct {
	Name      string
	Filename  string
	File      *ast.File
	Doc       Doc
	EventType EventType
	Source    interface{}
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
