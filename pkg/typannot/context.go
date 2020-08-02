package typannot

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Context of annotation
	Context struct {
		*typgo.Context
		ASTStore *ASTStore
		Dirs     []string
	}
)

// CreateContext return new instance of context
func CreateContext(c *typgo.Context) (*Context, error) {
	dirs, files := WalkLayout(c.BuildSys.ProjectLayouts)
	astStore, err := CreateASTStore(files...)
	if err != nil {
		return nil, err
	}
	return &Context{
		Context:  c,
		ASTStore: astStore,
		Dirs:     dirs,
	}, nil
}

// CreateImports create import line
func (c *Context) CreateImports(projPkg string, more ...string) []string {
	var imports []string
	for _, dir := range c.Dirs {
		imports = append(imports, fmt.Sprintf("%s/%s", projPkg, dir))
	}
	imports = append(imports, more...)
	return imports
}

// WalkLayout return dirs and files
func WalkLayout(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if isGoSource(path) {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

func isGoSource(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}

//
// Context
//

// FindAnnotByFunc find annotation by function
func (c *Context) FindAnnotByFunc(tagName string) []*Annot {
	var annots []*Annot
	for _, annot := range c.ASTStore.Annots {
		_, ok := annot.Type.(*FuncType)
		if strings.EqualFold(tagName, annot.TagName) && ok {
			annots = append(annots, annot)
		}
	}
	return annots
}

// FindAnnotByStruct find annotation by struct
func (c *Context) FindAnnotByStruct(tagName string) []*Annot {
	var annots []*Annot
	for _, annot := range c.ASTStore.Annots {
		_, ok := annot.Type.(*StructType)
		if strings.EqualFold(tagName, annot.TagName) && ok {
			annots = append(annots, annot)
		}
	}
	return annots
}

// FindAnnotByInterface find annotation by interface
func (c *Context) FindAnnotByInterface(tagName string) []*Annot {
	var annots []*Annot
	for _, annot := range c.ASTStore.Annots {
		_, ok := annot.Type.(*InterfaceType)
		if strings.EqualFold(tagName, annot.TagName) && ok {
			annots = append(annots, annot)
		}
	}
	return annots
}
