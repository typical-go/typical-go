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
	Docs      []*ast.CommentGroup
	DeclNodes []ast.Decl
	Annots    []*Annot
}

func (a *ASTStore) put(decl *Decl, node ast.Decl, doc *ast.CommentGroup) {
	a.Decls = append(a.Decls, decl)
	a.DeclNodes = append(a.DeclNodes, node)
	a.Docs = append(a.Docs, doc)
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
			putDecls(store, node, path, pkg)
		}
	}

	for i, decl := range store.Decls {
		if err = putAnnots(store, decl, store.Docs[i]); err != nil {
			return
		}
	}

	return store, nil
}

func putDecls(store *ASTStore, node ast.Decl, path, pkg string) {
	switch node.(type) {
	case *ast.FuncDecl:
		funcDecl := node.(*ast.FuncDecl)
		doc := funcDecl.Doc
		name := funcDecl.Name.Name

		store.put(&Decl{
			Name:    name,
			Type:    FuncType,
			Path:    path,
			Package: pkg,
		}, node, doc)

	case *ast.GenDecl:
		genDecl := node.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				typeSpec := spec.(*ast.TypeSpec)

				declType := GenericType
				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					declType = InterfaceType
				case *ast.StructType:
					declType = StructType
				}

				// NOTE: get type specific first before get the generic
				doc := typeSpec.Doc
				if doc == nil {
					doc = genDecl.Doc
				}

				store.put(&Decl{
					Name:    typeSpec.Name.Name,
					Type:    declType,
					Path:    path,
					Package: pkg,
				}, node, doc)
			}
		}
	}

}

func putAnnots(store *ASTStore, decl *Decl, doc *ast.CommentGroup) (err error) {
	var rawAnnots []string

	if doc == nil {
		return
	}

	RetrRawAnnots(&rawAnnots, doc.Text())

	for _, raw := range rawAnnots {
		if a, _ := CreateAnnot(decl, raw); a != nil {
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
