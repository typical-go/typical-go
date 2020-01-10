package walker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Walker responsible to walk the filenames
type Walker struct {
	filenames []string
}

// New return new constructor of walker
func New(filenames []string) *Walker {
	return &Walker{
		filenames: filenames,
	}
}

// Walk the source code to get autowire and automock
func (w *Walker) Walk() (events DeclEvents, err error) {
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range w.filenames {
		if isWalkTarget(filename) {
			var f *ast.File
			if f, err = parser.ParseFile(fset, filename, nil, parser.ParseComments); err != nil {
				return
			}
			for _, decl := range f.Decls {
				events.Append(w.declaration(filename, f, decl)...)
			}
		}
	}
	return
}

func (w *Walker) declaration(filename string, f *ast.File, decl ast.Decl) (events DeclEvents) {
	switch decl.(type) {
	case *ast.FuncDecl:
		var (
			doc      string
			funcDecl = decl.(*ast.FuncDecl)
		)
		if funcDecl.Doc != nil {
			doc = funcDecl.Doc.Text()
		}
		events = append(events, &DeclEvent{
			EventType:   FunctionType,
			SourceName:  funcDecl.Name.Name,
			SourceObj:   funcDecl,
			Filename:    filename,
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
					typeSpec  = spec.(*ast.TypeSpec)
					eventType = GenericType
				)
				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					eventType = InterfaceType
				case *ast.StructType:
					eventType = StructType
				}
				events = append(events, &DeclEvent{
					EventType:   eventType,
					SourceName:  typeSpec.Name.Name,
					SourceObj:   typeSpec,
					Filename:    filename,
					File:        f,
					Doc:         doc,
					Annotations: ParseAnnotations(doc),
				})
			}
		}
	}
	return
}

func isWalkTarget(filename string) bool { //  TODO: move out from walker package
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
