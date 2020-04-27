package typast

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Walk the source code to get autowire and automock
func Walk(filenames ...string) (a *ASTStore, err error) {
	var (
		decls []*Decl
	)

	if decls, err = parseFiles(filenames); err != nil {
		return
	}

	return &ASTStore{
		filenames: filenames,
		decls:     decls,
	}, nil
}

func parseFiles(filenames []string) (decls []*Decl, err error) {
	var f *ast.File
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range filenames {
		if f, err = parser.ParseFile(fset, filename, nil, parser.ParseComments); err != nil {
			return
		}
		for _, d := range f.Decls {
			if decl := parseDecl(filename, f, d); decl != nil {
				decls = append(decls, decl)
			}
		}
	}
	return
}

func parseDecl(path string, f *ast.File, decl ast.Decl) *Decl {
	switch decl.(type) {
	case *ast.FuncDecl:
		funcDecl := decl.(*ast.FuncDecl)
		return &Decl{
			Type:       FunctionType,
			SourceName: funcDecl.Name.Name,
			SourceObj:  funcDecl,
			Path:       path,
			File:       f,
			Doc:        funcDecl.Doc,
		}
	case *ast.GenDecl:
		genDecl := decl.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				typeSpec := spec.(*ast.TypeSpec)
				declType := GenericType
				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					declType = InterfaceType
				case *ast.StructType:
					declType = StructType
				}
				return &Decl{
					Type:       declType,
					SourceName: typeSpec.Name.Name,
					SourceObj:  typeSpec,
					Path:       path,
					File:       f,
					Doc:        genDecl.Doc,
				}
			}
		}
	}
	return nil
}
