package typast

import (
	"go/ast"
)

type (
	// InterfaceDecl interface declaration
	InterfaceDecl struct {
		TypeDecl
	}
	// StructDecl struct declaration
	StructDecl struct {
		TypeDecl
		Fields []*Field
	}
	// TypeDecl type declaration
	TypeDecl struct {
		GenDecl
		Name string
		Docs []string
	}
	// GenDecl generic declaration
	GenDecl struct {
		Docs []string
	}
)

func createGenDecl(genDecl *ast.GenDecl, file File) []DeclType {
	var declTypes []DeclType
	for _, spec := range genDecl.Specs {
		switch spec.(type) {
		case *ast.TypeSpec:
			typeSpec := spec.(*ast.TypeSpec)
			typeDecl := TypeDecl{
				GenDecl: GenDecl{Docs: docs(genDecl.Doc)},
				Name:    typeSpec.Name.Name,
				Docs:    docs(typeSpec.Doc),
			}

			switch typeSpec.Type.(type) {
			case *ast.InterfaceType:
				declTypes = append(declTypes, createInterfaceDecl(typeDecl))
			case *ast.StructType:
				declTypes = append(declTypes, createStructDecl(typeDecl, typeSpec.Type.(*ast.StructType)))
			}
		}
	}
	return declTypes
}

func createInterfaceDecl(typeDecl TypeDecl) *InterfaceDecl {
	return &InterfaceDecl{TypeDecl: typeDecl}
}

func createStructDecl(typeDecl TypeDecl, structType *ast.StructType) *StructDecl {
	return &StructDecl{
		Fields:   structFields(structType),
		TypeDecl: typeDecl,
	}
}

func structFields(s *ast.StructType) []*Field {
	var fields []*Field
	for _, f := range s.Fields.List {
		fields = append(fields, createField(f))
	}
	return fields
}

func docs(group *ast.CommentGroup) []string {
	if group == nil {
		return nil
	}
	var docs []string
	for _, comment := range group.List {
		docs = append(docs, comment.Text)
	}
	return docs
}

//
// TypeDecl
//

var _ DeclType = (*TypeDecl)(nil)

// GetName get name
func (t *TypeDecl) GetName() string {
	return t.Name
}

// GetDocs get doc
func (t *TypeDecl) GetDocs() []string {
	if t.Docs != nil {
		return t.Docs
	}
	return t.GenDecl.Docs
}
