package typannot

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Annotators is extra information from source code
	Annotators []Annotator
	// Annotator responsible to annotate
	Annotator interface {
		Annotate(*Context) error
	}
	// Context of annotation
	Context struct {
		*typgo.Context
		ASTStore *ASTStore
		Imports  []string
	}
)

var _ typgo.Action = (Annotators)(nil)

// Execute annotation
func (a Annotators) Execute(c *typgo.Context) error {
	annCtx, err := CreateContext(c)
	if err != nil {
		return err
	}

	for _, annotator := range a {
		if err := annotator.Annotate(annCtx); err != nil {
			return err
		}
	}

	return nil
}

// CreateContext return new instance of context
func CreateContext(c *typgo.Context) (*Context, error) {
	appDirs, appFiles := WalkLayout(c.Layouts)
	astStore, err := CreateASTStore(appFiles...)
	if err != nil {
		return nil, err
	}
	imports := retrImports(appDirs)
	return &Context{
		Context:  c,
		ASTStore: astStore,
		Imports:  imports,
	}, nil
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typapp",
	}

	for _, dir := range dirs {
		imports = append(imports, fmt.Sprintf("%s/%s", typgo.ProjectPkg, dir))
	}
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
