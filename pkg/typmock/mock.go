package typmock

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
)

// Mock annotation data
type Mock struct {
	Dir     string `json:"-"`
	Pkg     string `json:"-"`
	Source  string `json:"-"`
	Parent  string `json:"-"`
	MockPkg string `json:"-"`
}

// CreateMock to create mock
func CreateMock(annot *typast.Annot) *Mock {
	if !isMock(annot) {
		return nil
	}

	pkg := annot.Decl.Pkg
	dir := filepath.Dir(annot.Decl.Path)

	parent := ""
	if dir != "." {
		parent = dir[:len(dir)-len(pkg)]
	}

	return &Mock{
		Dir:     dir,
		Pkg:     pkg,
		Source:  annot.Decl.Name,
		Parent:  parent,
		MockPkg: fmt.Sprintf("%s_mock", pkg),
	}
}

func isMock(annot *typast.Annot) bool {
	return strings.EqualFold(annot.TagName, MockTag) &&
		annot.Decl.Type == typast.Interface
}
