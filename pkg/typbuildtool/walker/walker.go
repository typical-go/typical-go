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
func (w *Walker) AddFuncDeclListener(listener FuncDeclListener) *Walker {
	w.funcDeclListeners = append(w.funcDeclListeners, listener)
	return w
}

// AddTypeSpecListener to add function declaration listener
func (w *Walker) AddTypeSpecListener(listener TypeSpecListener) *Walker {
	w.typeSpecListeners = append(w.typeSpecListeners, listener)
	return w
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
	for _, decl := range f.Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl.(*ast.FuncDecl)
			e := &FuncDeclEvent{
				Name:     funcDecl.Name.Name,
				Filename: filename,
				File:     f,
				FuncDecl: funcDecl,
			}
			for _, listener := range w.funcDeclListeners {
				if err = listener.OnFuncDecl(e); err != nil {
					return
				}
			}
		case *ast.GenDecl:
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					e := &TypeSpecEvent{
						Name:     typeSpec.Name.Name,
						Filename: filename,
						File:     f,
						TypeSpec: typeSpec,
						GenDecl:  genDecl,
					}
					for _, listener := range w.typeSpecListeners {
						if err = listener.OnTypeSpec(e); err != nil {
							return
						}
					}
				}
			}

		}
	}
	return
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
