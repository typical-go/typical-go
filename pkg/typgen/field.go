package typgen

import (
	"fmt"
	"go/ast"
	"reflect"
)

type (
	// Field information
	Field struct {
		Names []string
		Type  string
		reflect.StructTag
	}
	// FieldList function parameter
	FieldList struct {
		List []*Field
	}
)

func createField(f *ast.Field) *Field {
	var names []string
	var typ string

	structTag := StructTag(f.Tag)

	for _, n := range f.Names {
		names = append(names, n.Name)
	}

	switch f.Type.(type) {
	case *ast.Ident:
		typ = f.Type.(*ast.Ident).Name
	case *ast.StarExpr:
		typ = fmt.Sprintf("*%s", f.Type.(*ast.StarExpr).X)
	}
	return &Field{
		Names:     names,
		Type:      typ,
		StructTag: structTag,
	}
}

// func createFuncParam(l *ast.FieldList) *FieldList {
// 	var list []*Field
// 	if l != nil {
// 		for _, f := range l.List {
// 			list = append(list, createField(f))
// 		}
// 	}
// 	return &FieldList{List: list}
// }

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
