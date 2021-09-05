package typgen

import (
	"go/ast"
)

type (
	// StructDecl struct declaration
	StructDecl struct {
		TypeDecl
		Fields []*Field
	}
)

func CreateStructDecl(typeDecl TypeDecl, structType *ast.StructType) *StructDecl {
	var fields []*Field
	for _, f := range structType.Fields.List {
		fields = append(fields, createField(f))
	}
	return &StructDecl{
		Fields:   fields,
		TypeDecl: typeDecl,
	}
}
