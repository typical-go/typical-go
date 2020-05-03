package typannot

import (
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	// MockTag is tag for mock
	MockTag = "mock"
)

// Mock annotation data
type Mock struct {
	Dir    string `json:"-"`
	Pkg    string `json:"-"`
	Source string `json:"-"`
	Parent string `json:"-"`
}

// GetMock to get mock annotation from ast store
func GetMock(store *typast.ASTStore) (mocks []*Mock) {
	for _, annot := range store.Annots {
		if isMock(annot) {
			pkg := annot.Pkg
			dir := filepath.Dir(annot.Path)

			parent := ""
			if dir != "." {
				parent = dir[:len(dir)-len(pkg)]
			}

			mocks = append(mocks, &Mock{
				Dir:    dir,
				Pkg:    pkg,
				Source: annot.Name,
				Parent: parent,
			})
		}
	}
	return
}

func isMock(annot *typast.Annotation) bool {
	return strings.EqualFold(annot.TagName, MockTag) &&
		annot.Type == typast.Interface
}
