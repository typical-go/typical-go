package typgen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

func Compile(paths ...string) ([]*Annotation, error) {
	var directives []*Annotation
	fset := token.NewFileSet() // positions are relative to fset

	for _, path := range paths {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		file := &File{
			Path:   path,
			Name:   f.Name.Name,
			Import: CreateImports(f),
		}

		for _, decl := range f.Decls {
			switch decl.(type) {
			case *ast.FuncDecl:
				declType := CreateFuncDecl(decl.(*ast.FuncDecl), file)
				directives = appendDecl(directives, file, declType)
			case *ast.GenDecl:
				declTypes := createGenDecl(decl.(*ast.GenDecl), file)
				for _, declType := range declTypes {
					directives = appendDecl(directives, file, declType)
				}
			}
		}
	}

	return directives, nil
}

func appendDecl(d []*Annotation, file *File, declType Type) []*Annotation {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	return append(d, retrieveAnnots(decl)...)
}

func retrieveAnnots(decl *Decl) []*Annotation {
	var annots []*Annotation
	for _, raw := range decl.GetDocs() {
		if strings.HasPrefix(raw, "//") {
			raw = strings.TrimSpace(raw[2:])
		}
		if strings.HasPrefix(raw, "@") {
			name, params := ParseRawAnnot(raw)
			annots = append(annots, &Annotation{
				Name:   name,
				Params: reflect.StructTag(params),
				Decl:   decl,
			})
		}
	}

	return annots
}

// ParseRawAnnot parse raw string to annotation
func ParseRawAnnot(raw string) (name, params string) {
	iOpen := strings.IndexRune(raw, '(')
	iSpace := strings.IndexRune(raw, ' ')

	if iOpen < 0 {
		if iSpace < 0 {
			return strings.TrimSpace(raw), ""
		}
		name = raw[:iSpace]
	} else {
		if iSpace < 0 {
			name = raw[:iOpen]
		} else {
			name = raw[:iSpace]
		}

		if iClose := strings.IndexRune(raw, ')'); iClose > 0 {
			params = raw[iOpen+1 : iClose]
		}
	}

	return name, params
}
