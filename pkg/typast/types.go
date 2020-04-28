package typast

import (
	"go/ast"
)

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
	Type       DeclType
	Doc        *ast.CommentGroup
	SourceName string
	SourceObj  interface{}
}

// DeclType is declaration type
type DeclType string
