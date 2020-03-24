package typast

import (
	"go/ast"
)

// DeclType is declaration type
type DeclType string

const (
	FunctionType  = DeclType("Function")
	InterfaceType = DeclType("Interface")
	StructType    = DeclType("Struct")
	GenericType   = DeclType("Generic")
)

// Declaration information
type Declaration struct {
	Path        string
	File        *ast.File
	Doc         string
	Annotations []*Annotation
	Type        DeclType
	SourceName  string
	SourceObj   interface{}
}

// DeclFunc to handle declaration
type DeclFunc func(*Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *Declaration, ann *Annotation) error
