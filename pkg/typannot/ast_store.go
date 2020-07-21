package typannot

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type (
	// ASTStore responsible to store filename, declaration and annotation
	ASTStore struct {
		Paths  []string
		Decls  []*Decl
		Annots []*Annot
	}
	// Decl stand of declaration
	Decl struct {
		Name    string
		Path    string
		Package string
		Type    interface{}
	}
	// FuncType function type
	FuncType struct{}
	// InterfaceType interface type
	InterfaceType struct{}
	// StructType struct type
	StructType struct{}
)

// CreateASTStore to walk through the filenames and store declaration and annotations
func CreateASTStore(paths ...string) (*ASTStore, error) {
	var (
		decls  []*Decl
		annots []*Annot
	)

	fset := token.NewFileSet() // positions are relative to fset
	for _, path := range paths {

		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		pkg := f.Name.Name
		for _, node := range f.Decls {
			switch node.(type) {
			case *ast.FuncDecl:
				funcDecl := node.(*ast.FuncDecl)
				name := funcDecl.Name.Name
				decl := &Decl{Name: name, Type: &FuncType{}, Path: path, Package: pkg}

				decls = append(decls, decl)
				annots = append(annots, retrieveAnnots(decl, funcDecl.Doc)...)

			case *ast.GenDecl:
				genDecl := node.(*ast.GenDecl)
				for _, spec := range genDecl.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typeSpec := spec.(*ast.TypeSpec)

						var declType interface{}
						switch typeSpec.Type.(type) {
						case *ast.InterfaceType:
							declType = &InterfaceType{}
						case *ast.StructType:
							declType = &StructType{}
						}

						// NOTE: get type specific first before get the generic
						doc := typeSpec.Doc
						if doc == nil {
							doc = genDecl.Doc
						}

						name := typeSpec.Name.Name
						decl := &Decl{Name: name, Type: declType, Path: path, Package: pkg}

						decls = append(decls, decl)
						annots = append(annots, retrieveAnnots(decl, doc)...)
					}
				}
			}
		}
	}

	return &ASTStore{Paths: paths, Decls: decls, Annots: annots}, nil
}

func retrieveAnnots(decl *Decl, doc *ast.CommentGroup) []*Annot {
	var rawAnnots []string
	RetrieveRawAnnots(&rawAnnots, doc.Text())

	var annots []*Annot
	for _, raw := range rawAnnots {
		if annot, _ := CreateAnnot(decl, raw); annot != nil {
			annots = append(annots, annot)
		}
	}

	return annots
}

// RetrieveRawAnnots retrieve raw of annotation for godoc text
func RetrieveRawAnnots(rawAnnots *[]string, docText string) {
	docText = strings.TrimSpace(docText)
	enter := strings.IndexRune(docText, '\n')
	if !strings.HasPrefix(docText, "@") {
		if enter > 0 {
			RetrieveRawAnnots(rawAnnots, docText[enter+1:])
		}
		return
	}

	open := strings.IndexRune(docText, '{')

	if enter < open && enter > 0 {
		*rawAnnots = append(*rawAnnots, docText[:enter])
		RetrieveRawAnnots(rawAnnots, docText[enter+1:])
		return
	}

	if open < 0 {
		if enter < 0 {
			*rawAnnots = append(*rawAnnots, docText)
		} else {
			*rawAnnots = append(*rawAnnots, docText[:enter])
			RetrieveRawAnnots(rawAnnots, docText[enter+1:])
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
	RetrieveRawAnnots(rawAnnots, docText[enter2+1:])
	return
}
