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
	typeSpecListeners []TypeSpecListener
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

// AddTypeSpecListener to add function declaration listener
func (w *Walker) AddTypeSpecListener(listener TypeSpecListener) {
	w.typeSpecListeners = append(w.typeSpecListeners, listener)
}

// Walk the source code to get autowire and automock
func (w *Walker) Walk() (err error) {
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range w.filenames {
		if isWalkTarget(filename) {
			if err = w.parse(fset, filename); err != nil {
				return
			}
		}
	}
	return
}

func (w *Walker) parse(fset *token.FileSet, filename string) (err error) {
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return
	}
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
			e := &TypeSpecEvent{
				Name:     name,
				File:     f,
				Filename: filename,
				TypeSpec: obj.Decl.(*ast.TypeSpec),
			}
			for _, listener := range w.typeSpecListeners {
				listener.OnTypeSpec(e)
			}
		}
	}
	return
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
