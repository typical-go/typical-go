package walker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Walker responsible to walk the filenames
type Walker struct {
	filenames     []string
	declListeners []DeclListener
}

// New return new constructor of walker
func New(filenames []string) *Walker {
	return &Walker{
		filenames: filenames,
	}
}

// AddDeclListener to add function declaration listener
func (w *Walker) AddDeclListener(listener DeclListener) *Walker {
	w.declListeners = append(w.declListeners, listener)
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
			var (
				doc      string
				funcDecl = decl.(*ast.FuncDecl)
			)
			if funcDecl.Doc != nil {
				doc = funcDecl.Doc.Text()
			}
			if err = w.fireEvent(&DeclEvent{
				Name:      funcDecl.Name.Name,
				Filename:  filename,
				File:      f,
				Doc:       Doc(doc),
				EventType: FunctionType,
				Source:    funcDecl,
			}); err != nil {
				return
			}
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
					if err = w.fireEvent(&DeclEvent{
						Name:      typeSpec.Name.Name,
						Filename:  filename,
						File:      f,
						Doc:       Doc(doc),
						EventType: eventType,
						Source:    typeSpec,
					}); err != nil {
						return
					}
				}
			}

		}
	}
	return
}

func (w *Walker) fireEvent(e *DeclEvent) (err error) {
	for _, listener := range w.declListeners {
		if err = listener.OnDecl(e); err != nil {
			return
		}
	}
	return
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
