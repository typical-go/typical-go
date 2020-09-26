package typapp

import (
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Annot struct {
		*typast.Annot
		ImportAlias string
	}
)

// FindAnnotFunc ...
func FindAnnotFunc(c *typast.Context, tagName string) ([]*Annot, map[string]string) {
	var annots []*Annot
	imports := make(map[string]string)
	lastAlias := ""

	for _, annot := range c.FindAnnot(tagName, typast.EqualFunc) {
		importPath := fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(annot.Path))
		importAlias, ok := imports[importPath]
		if !ok {
			lastAlias = nextAlias(lastAlias)
			imports[importPath] = lastAlias
			importAlias = lastAlias
		}
		annots = append(annots, &Annot{
			Annot:       annot,
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
