package typast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// ASTStore responsible to store filename, declaration and annotation
type ASTStore struct {
	Paths     []string
	Decls     []*Decl
	DeclNodes []ast.Decl
	Annots    []*Annotation
}

// CreateASTStore to walk through the filenames and store declaration and annotations
func CreateASTStore(paths ...string) *ASTStore {
	var (
		decls     []*Decl
		declNodes []ast.Decl
		annots    []*Annotation
		err       error
	)

	fset := token.NewFileSet() // positions are relative to fset
	for _, path := range paths {
		var f *ast.File
		if f, err = parser.ParseFile(fset, path, nil, parser.ParseComments); err != nil {
			panic(err.Error())
		}

		pkg := f.Name.Name
		for _, node := range f.Decls {
			name, declType, doc := parseDecl(node)
			if name != "" {
				declNodes = append(declNodes, node)
				decl := &Decl{
					Name: name,
					Type: declType,
					Path: path,
					Pkg:  pkg,
				}
				decls = append(decls, decl)
				annots = append(annots, parseAnnots(decl, doc)...)
			}
		}
	}

	return &ASTStore{
		Paths:     paths,
		Decls:     decls,
		DeclNodes: declNodes,
		Annots:    annots,
	}
}

func parseDecl(decl ast.Decl) (name string, declType DeclType, doc *ast.CommentGroup) {
	switch decl.(type) {
	case *ast.FuncDecl:
		funcDecl := decl.(*ast.FuncDecl)
		name = funcDecl.Name.Name
		declType = Function
		doc = funcDecl.Doc
		return
	case *ast.GenDecl:
		genDecl := decl.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				typeSpec := spec.(*ast.TypeSpec)
				name = typeSpec.Name.Name
				doc = genDecl.Doc

				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					declType = Interface
				case *ast.StructType:
					declType = Struct
				default:
					declType = Generic
				}
				return
			}
		}
	}
	return
}

func parseAnnots(decl *Decl, doc *ast.CommentGroup) (annotations []*Annotation) {
	if doc == nil {
		return
	}

	for _, line := range strings.Split(doc.Text(), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "@") {
			a := CreateAnnotation(decl, line)
			if a != nil {
				annotations = append(annotations, a)
			}
		}
	}

	return
}
