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

// Send event to listeners
func (e *DeclEvent) Send(listeners ...DeclListener) (err error) {
	for _, listener := range listeners {
		if err = listener.OnDecl(e); err != nil {
			return
		}
	}
	return
}

type DeclEvents []*DeclEvent

// Append event
func (d *DeclEvents) Append(e ...*DeclEvent) *DeclEvents {
	*d = append(*d, e...)
	return d
}

// Send events to listeners
func (d *DeclEvents) Send(listeners ...DeclListener) (err error) {
	for _, event := range *d {
		event.Send(listeners...)
	}
	return
}
