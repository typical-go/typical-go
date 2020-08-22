package typannot

import (
	"go/ast"
	"reflect"
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
	// Field information
	Field struct {
		Name string
		Type string
		reflect.StructTag
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
	for _, field := range s.Fields.List {
		switch field.Type.(type) {
		case *ast.Ident:
			i := field.Type.(*ast.Ident)
			for _, name := range field.Names {
				fields = append(fields, &Field{
					Name:      name.Name,
					Type:      i.Name,
					StructTag: StructTag(field.Tag),
				})
			}
		}
	}
	return fields
}

// StructTag create struct tag
func StructTag(tag *ast.BasicLit) reflect.StructTag {
	if tag == nil {
		return ""
	}
	s := tag.Value
	n := len(s)
	if n < 2 {
		return ""
	}
	return reflect.StructTag(s[1 : n-1])
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
