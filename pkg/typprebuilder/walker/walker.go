package walker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Walker responsible to walk the filenames
type Walker struct {
	filenames         []string
	funcDeclListeners []FuncDeclListener
}

// New return new constructor of walker
func New(filenames []string) *Walker {
	return &Walker{
		filenames: filenames,
	}
}

// AddFuncDeclListener to add function declaration listener
func (w *Walker) AddFuncDeclListener(listener FuncDeclListener) {
	w.funcDeclListeners = append(w.funcDeclListeners, listener)
}

// Walk the source code to get autowire and automock
func (w *Walker) Walk() (files *ProjectFiles, err error) {
	files = &ProjectFiles{}
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range w.filenames {
		if isWalkTarget(filename) {
			var file ProjectFile
			if file, err = w.parse(fset, filename); err != nil {
				return
			}
			if !file.IsEmpty() {
				files.Add(file)
			}
		}
	}
	return
}

func (w *Walker) parse(fset *token.FileSet, filename string) (projFile ProjectFile, err error) {
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return
	}
	projFile.Name = filename
	for name, obj := range f.Scope.Objects {
		switch obj.Decl.(type) {
		case *ast.FuncDecl:
			e := &FuncDeclEvent{
				Name:     name,
				File:     f,
				FuncDecl: obj.Decl.(*ast.FuncDecl),
			}
			for _, listener := range w.funcDeclListeners {
				if listener.IsAction(e) {
					if err = listener.ActionPerformed(e); err != nil {
						return
					}
				}
			}
		case *ast.TypeSpec:
			typeSpec := obj.Decl.(*ast.TypeSpec)
			switch typeSpec.Type.(type) {
			case *ast.StructType:
			case *ast.InterfaceType:
				var doc string
				if typeSpec.Doc != nil {
					doc = typeSpec.Doc.Text()
				}
				projFile.Mock = isAutoMock(doc)
			}
		}
	}
	return
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}

func isAutoMock(doc string) bool {
	notations := ParseAnnotations(doc)
	return !notations.Contain("nomock")
}
