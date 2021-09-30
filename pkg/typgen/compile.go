package typgen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

func Compile(paths ...string) ([]*Directive, error) {
	var directives []*Directive
	fset := token.NewFileSet() // positions are relative to fset

	for _, path := range paths {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		file := CreateFile(path, f)

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

func appendDecl(d []*Directive, file *File, declType Type) []*Directive {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	// s.Decls = append(s.Decls, decl)
	return append(d, retrieveAnnots(decl)...)
}

func retrieveAnnots(decl *Decl) []*Directive {
	var annots []*Directive
	for _, raw := range decl.GetDocs() {
		if strings.HasPrefix(raw, "//") {
			raw = strings.TrimSpace(raw[2:])
		}
		if strings.HasPrefix(raw, "@") {
			tagName, tagAttrs := ParseRawAnnot(raw)
			annots = append(annots, &Directive{
				TagName:  tagName,
				TagParam: reflect.StructTag(tagAttrs),
				Decl:     decl,
			})
		}
	}

	return annots
}

// ParseRawAnnot parse raw string to annotation
func ParseRawAnnot(raw string) (tagName, tagAttrs string) {
	iOpen := strings.IndexRune(raw, '(')
	iSpace := strings.IndexRune(raw, ' ')

	if iOpen < 0 {
		if iSpace < 0 {
			tagName = strings.TrimSpace(raw)
			return tagName, ""
		}
		tagName = raw[:iSpace]
	} else {
		if iSpace < 0 {
			tagName = raw[:iOpen]
		} else {
			tagName = raw[:iSpace]
		}

		if iClose := strings.IndexRune(raw, ')'); iClose > 0 {
			tagAttrs = raw[iOpen+1 : iClose]
		}
	}

	return tagName, tagAttrs
}
