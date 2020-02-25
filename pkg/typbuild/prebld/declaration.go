package prebld

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
	Annotations Annotations
	Type        DeclType
	SourceName  string
	SourceObj   interface{}
}
