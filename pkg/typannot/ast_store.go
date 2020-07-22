package typannot

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type (
	// ASTStore responsible to store filename, declaration and annotation
	ASTStore struct {
		Paths  []string
		Decls  []*Decl
		Annots []*Annot
	}
	// Decl stand of declaration
	Decl struct {
		Name    string
		Path    string
		Package string
		Type    interface{}
	}
	// FuncType function type
	FuncType struct{}
	// InterfaceType interface type
	InterfaceType struct{}
	// StructType struct type
	StructType struct {
		Fields []*Field
	}
	// Field information
	Field struct {
		Name string
		Type string
		reflect.StructTag
	}
)

// CreateASTStore to walk through the filenames and store declaration and annotations
func CreateASTStore(paths ...string) (*ASTStore, error) {
	var (
		decls  []*Decl
		annots []*Annot
	)

	fset := token.NewFileSet() // positions are relative to fset
	for _, path := range paths {

		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		pkg := f.Name.Name
		for _, node := range f.Decls {
			switch node.(type) {
			case *ast.FuncDecl:
				funcDecl := node.(*ast.FuncDecl)
				name := funcDecl.Name.Name
				decl := &Decl{Name: name, Type: &FuncType{}, Path: path, Package: pkg}

				decls = append(decls, decl)
				annots = append(annots, retrieveAnnots(decl, funcDecl.Doc)...)

			case *ast.GenDecl:
				genDecl := node.(*ast.GenDecl)
				for _, spec := range genDecl.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typeSpec := spec.(*ast.TypeSpec)

						var typ interface{}
						switch typeSpec.Type.(type) {
						case *ast.InterfaceType:
							typ = &InterfaceType{}
						case *ast.StructType:
							typ = convertStructType(typeSpec.Type.(*ast.StructType))
						}

						// NOTE: get type specific first before get the generic
						doc := typeSpec.Doc
						if doc == nil {
							doc = genDecl.Doc
						}

						name := typeSpec.Name.Name
						decl := &Decl{Name: name, Type: typ, Path: path, Package: pkg}

						decls = append(decls, decl)
						annots = append(annots, retrieveAnnots(decl, doc)...)
					}
				}
			}
		}
	}

	return &ASTStore{Paths: paths, Decls: decls, Annots: annots}, nil
}

func convertStructType(s *ast.StructType) *StructType {
	var fields []*Field
	for _, field := range s.Fields.List {
		switch field.Type.(type) {
		case *ast.Ident:
			i := field.Type.(*ast.Ident)
			for _, name := range field.Names {
				fields = append(fields, &Field{
					Name:      name.Name,
					Type:      i.Name,
					StructTag: nakedStructTag(field.Tag.Value),
				})
			}
		}
	}
	return &StructType{Fields: fields}
}

func nakedStructTag(s string) reflect.StructTag {
	n := len(s)
	if n < 2 {
		return ""
	}
	return reflect.StructTag(s[1 : len(s)-1])
}

func retrieveAnnots(decl *Decl, doc *ast.CommentGroup) []*Annot {
	if doc == nil {
		return nil
	}

	var annots []*Annot
	for _, comment := range doc.List {
		raw := comment.Text
		if strings.HasPrefix(raw, "//") {
			raw = strings.TrimSpace(raw[2:])
		}
		if strings.HasPrefix(raw, "@") {
			tagName, tagAttrs := ParseAnnot(raw)
			annots = append(annots, &Annot{
				TagName:  tagName,
				TagAttrs: reflect.StructTag(tagAttrs),
				Decl:     decl,
			})
		}
	}

	return annots
}
