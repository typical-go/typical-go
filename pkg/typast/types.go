package typast

import (
	"go/ast"
)

const (
	// Function type
	Function DeclType = iota

	// Interface type
	Interface

	// Struct type
	Struct

	// Generic type
	Generic
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
type DeclType int

func (d DeclType) String() string {
	return [...]string{"Function", "Interface", "Struct", "Generic"}[d]
}
