package typast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type (
	// Summary responsible to store filename, declaration and annotation
	Summary struct {
		Paths  []string
		Decls  []*Decl
		Annots []*Directive
	}
	// Decl stand of declaration
	Decl struct {
		File
		Type
	}
	// Type declaratio type
	Type interface {
		GetName() string
		GetDocs() []string
	}
	// File information
	File struct {
		Path    string
		Package string
	}
)

// Compile paths to ASTStore
func Compile(paths ...string) (*Summary, error) {
	summary := &Summary{Paths: paths}
	fset := token.NewFileSet() // positions are relative to fset

	for _, path := range paths {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		file := File{Path: path, Package: f.Name.Name}
		for _, decl := range f.Decls {
			switch decl.(type) {
			case *ast.FuncDecl:
				declType := createFuncDecl(decl.(*ast.FuncDecl), file)
				summary.AddDecl(file, declType)
			case *ast.GenDecl:
				declTypes := createGenDecl(decl.(*ast.GenDecl), file)
				for _, declType := range declTypes {
					summary.AddDecl(file, declType)
				}
			}
		}
	}

	return summary, nil
}

// Walk return dirs and files
func Walk(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

//
// Summary
//

// AddDecl add declaration
func (s *Summary) AddDecl(file File, declType Type) {
	decl := &Decl{
		File: file,
		Type: declType,
	}
	s.Decls = append(s.Decls, decl)
	s.Annots = append(s.Annots, retrieveAnnots(decl)...)
}
