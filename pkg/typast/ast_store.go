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
func CreateASTStore(paths ...string) (store *ASTStore, err error) {

	store = &ASTStore{
		Paths: paths,
	}

	fset := token.NewFileSet() // positions are relative to fset
	for _, path := range paths {
		var f *ast.File
		if f, err = parser.ParseFile(fset, path, nil, parser.ParseComments); err != nil {
			return
		}

		pkg := f.Name.Name
		for _, node := range f.Decls {
			store.DeclNodes = append(store.DeclNodes, node)
			name, declType, doc := parseDecl(node)
			decl := &Decl{
				Name: name,
				Type: declType,
				Path: path,
				Pkg:  pkg,
			}
			store.Decls = append(store.Decls, decl)

			if err = putAnnots(store, decl, doc); err != nil {
				return
			}

		}
	}

	return store, nil
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

func putAnnots(store *ASTStore, decl *Decl, doc *ast.CommentGroup) (err error) {
	var rawAnnots []string

	if doc == nil {
		return
	}

	RetrRawAnnots(&rawAnnots, doc.Text())

	for _, line := range rawAnnots {
		var a *Annotation
		a, _ = CreateAnnotation(decl, line)
		if a != nil {
			store.Annots = append(store.Annots, a)
		}
	}

	return
}

// RetrRawAnnots retrieve raw of annotation for godoc text
func RetrRawAnnots(rawAnnots *[]string, docText string) {
	docText = strings.TrimSpace(docText)
	enter := strings.IndexRune(docText, '\n')
	if !strings.HasPrefix(docText, "@") {
		if enter > 0 {
			RetrRawAnnots(rawAnnots, docText[enter+1:])
		}
		return
	}

	open := strings.IndexRune(docText, '{')

	if enter < open && enter > 0 {
		*rawAnnots = append(*rawAnnots, docText[:enter])
		RetrRawAnnots(rawAnnots, docText[enter+1:])
		return
	}

	if open < 0 {
		if enter < 0 {
			*rawAnnots = append(*rawAnnots, docText)
		} else {
			*rawAnnots = append(*rawAnnots, docText[:enter])
			RetrRawAnnots(rawAnnots, docText[enter+1:])
		}
		return
	}

	close := strings.IndexRune(docText, '}')
	if close < 0 {
		*rawAnnots = append(*rawAnnots, docText)
		return
	}

	enter2 := strings.IndexRune(docText[close:], '\n')
	if enter2 < 0 {
		*rawAnnots = append(*rawAnnots, docText)
		return
	}
	enter2 = close + enter2
	*rawAnnots = append(*rawAnnots, docText[:enter2])
	RetrRawAnnots(rawAnnots, docText[enter2+1:])
	return
}
