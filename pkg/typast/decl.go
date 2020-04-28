package typast

import (
	"fmt"
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

// Equal return true the declaration has same signature
func (d Decl) Equal(d1 *Decl) bool {
	return d1.SourceName == d.SourceName &&
		d1.Path == d.Path &&
		d1.Type == d.Type
}

func (d Decl) String() string {
	return fmt.Sprintf("Path:%s\tType: %s\tSourceName: %s",
		d.Path, d.Type, d.SourceName)
}

// DeclType is declaration type
type DeclType int

func (d DeclType) String() string {
	return [...]string{"Function", "Interface", "Struct", "Generic"}[d]
}
