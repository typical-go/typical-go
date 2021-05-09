package typgen

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Generator struct {
		Walker    Walker
		Processor Processor
	}
)

//
// Generator
//

var _ typgo.Tasker = (*Generator)(nil)
var _ typgo.Action = (*Generator)(nil)

// Task to annotate
func (a *Generator) Task() *typgo.Task {
	return &typgo.Task{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate code based on annotation directive ('@')",
		Action:  a,
	}
}

// Execute annotation
func (a *Generator) Execute(c *typgo.Context) error {
	filePaths := a.walk()
	if len(filePaths) < 1 {
		return errors.New("walker couldn't find any filepath")
	}
	directives, err := compile(filePaths...)
	if err != nil {
		return err
	}
	return a.Processor.Process(c, directives)
}

func (a *Generator) walk() []string {
	if a.Walker == nil {
		return Layouts{"internal"}.Walk()
	}
	return a.Walker.Walk()
}

func compile(paths ...string) (Directives, error) {
	var directives Directives
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
				directives.AddDecl(file, declType)
			case *ast.GenDecl:
				declTypes := createGenDecl(decl.(*ast.GenDecl), file)
				for _, declType := range declTypes {
					directives.AddDecl(file, declType)
				}
			}
		}
	}

	return directives, nil
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
