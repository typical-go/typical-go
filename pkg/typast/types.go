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

// Decl stand of declaration
type Decl struct {
	Path       string
	File       *ast.File
	Doc        string
	Annots     []*Annot
	Type       DeclType
	SourceName string
	SourceObj  interface{}
}

// DeclFunc to handle declaration
type DeclFunc func(*Decl) error

// AnnotFunc to handle annotation
type AnnotFunc func(decl *Decl, ann *Annot) error
