package typbuildtool

import (
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typast"
)

// MockTarget to be mocked
type MockTarget struct {
	SrcDir  string
	SrcPkg  string
	SrcName string
	MockPkg string
	MockDir string
	Dest    string
}

func createMockTarget(c *Context, decl *typast.Declaration) *MockTarget {
	var (
		pkg     = decl.File.Name.Name
		dir     = filepath.Dir(decl.Path)
		dirDest = dir[:len(dir)-len(pkg)]
		srcPkg  = fmt.Sprintf("%s/%s", c.ProjectPackage, dir)
		mockPkg = fmt.Sprintf("mock_%s", pkg)
		mockDir = fmt.Sprintf("%s%s", dirDest, mockPkg)
		dest    = fmt.Sprintf("%s/%s.go", mockDir, strcase.ToSnake(decl.SourceName))
	)
	return &MockTarget{
		SrcPkg:  srcPkg,
		SrcName: decl.SourceName,
		MockPkg: mockPkg,
		MockDir: mockDir,
		Dest:    dest,
	}
}
