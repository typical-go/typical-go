package typast

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Annot2 annotation with import alias
	Annot2 struct {
		*Annot
		Import      string
		ImportAlias string
	}
)

var _ CheckFn = EqualFunc
var _ CheckFn = EqualInterface
var _ CheckFn = EqualStruct

// EqualFunc return true if annot have tagName and public function
func EqualFunc(annot *Annot, tagName string) bool {
	funcDecl, ok := annot.Type.(*FuncDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) &&
		IsPublic(annot) && !funcDecl.IsMethod()
}

// EqualInterface return true if annot have tagName and public interface
func EqualInterface(annot *Annot, tagName string) bool {
	_, ok := annot.Type.(*InterfaceDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) && IsPublic(annot)
}

// EqualStruct return true if annot have tagName and public interface
func EqualStruct(annot *Annot, tagName string) bool {
	_, ok := annot.Type.(*StructDecl)
	return ok && strings.EqualFold(annot.TagName, tagName) && IsPublic(annot)
}

// IsPublic return true if decl is public access
func IsPublic(typ Type) bool {
	rune, _ := utf8.DecodeRuneInString(typ.GetName())
	return unicode.IsUpper(rune)
}

// FindAnnot ...
func FindAnnot(c *Context, tagName string, checkFn CheckFn) ([]*Annot2, map[string]string) {
	var annots []*Annot2
	imports := make(map[string]string)
	lastAlias := ""

	for _, annot := range c.FindAnnot(tagName, checkFn) {
		importPath := fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(annot.Path))
		importAlias, ok := imports[importPath]
		if !ok {
			lastAlias = nextAlias(lastAlias)
			imports[importPath] = lastAlias
			importAlias = lastAlias
		}
		annots = append(annots, &Annot2{
			Annot:       annot,
			Import:      importPath,
			ImportAlias: importAlias,
		})
	}
	return annots, imports
}

func nextAlias(last string) string {
	if last == "" {
		return "a"
	} else if last[len(last)-1] == 'z' {
		return last[:len(last)-1] + "aa"
	} else {
		return last[:len(last)-1] + string(last[len(last)-1]+1)
	}
}
