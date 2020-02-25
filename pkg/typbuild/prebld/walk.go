package prebld

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Walk the source code to get autowire and automock
func Walk(filenames []string) (store *DeclStore, err error) {
	store = &DeclStore{}
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range filenames {
		var f *ast.File
		if f, err = parser.ParseFile(fset, filename, nil, parser.ParseComments); err != nil {
			return
		}
		for _, decl := range f.Decls {
			store.Append(retrieveDeclarations(filename, f, decl)...)
		}
	}
	return
}

func retrieveDeclarations(path string, f *ast.File, decl ast.Decl) (declarations []*Declaration) {
	switch decl.(type) {
	case *ast.FuncDecl:
		var (
			doc      string
			funcDecl = decl.(*ast.FuncDecl)
		)
		if funcDecl.Doc != nil {
			doc = funcDecl.Doc.Text()
		}
		declarations = append(declarations, &Declaration{
			Type:        FunctionType,
			SourceName:  funcDecl.Name.Name,
			SourceObj:   funcDecl,
			Path:        path,
			File:        f,
			Annotations: ParseAnnotations(doc),
		})

	case *ast.GenDecl:
		var (
			doc     string
			genDecl = decl.(*ast.GenDecl)
		)
		if genDecl.Doc != nil {
			doc = genDecl.Doc.Text()
		}
		for _, spec := range genDecl.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				var (
					typeSpec = spec.(*ast.TypeSpec)
					declType = GenericType
				)
				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					declType = InterfaceType
				case *ast.StructType:
					declType = StructType
				}
				declarations = append(declarations, &Declaration{
					Type:        declType,
					SourceName:  typeSpec.Name.Name,
					SourceObj:   typeSpec,
					Path:        path,
					File:        f,
					Doc:         doc,
					Annotations: ParseAnnotations(doc),
				})
			}
		}
	}
	return
}
