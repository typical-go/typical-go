package typgen

import (
	"go/ast"
)

type (
	// Struct declaration
	Struct struct {
		TypeDecl
		Fields []*Field
	}
)

func CreateStructDecl(typeDecl TypeDecl, structType *ast.StructType) *Struct {
	var fields []*Field
	for _, f := range structType.Fields.List {
		fields = append(fields, createField(f))
	}
	return &Struct{
		Fields:   fields,
		TypeDecl: typeDecl,
	}
}
