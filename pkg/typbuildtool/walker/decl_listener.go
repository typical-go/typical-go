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
	Filename    string
	File        *ast.File
	Doc         string
	Annotations Annotations
	EventType   EventType
	SourceName  string
	SourceObj   interface{}
}
