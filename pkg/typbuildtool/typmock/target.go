package typmock

import (
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typast"
)

// Target to be mocked
type Target struct {
	SrcDir  string
	SrcPkg  string
	SrcName string
	MockPkg string
	MockDir string
	Dest    string
}

func createTarget(c *Context, decl *typast.Declaration) *Target {
	var (
		pkg     = decl.File.Name.Name
		dir     = filepath.Dir(decl.Path)
		dirDest = dir[:len(dir)-len(pkg)]
		srcPkg  = fmt.Sprintf("%s/%s", c.ModulePackage, dir)
		mockPkg = fmt.Sprintf("mock_%s", pkg)
		mockDir = fmt.Sprintf("%s%s", dirDest, mockPkg)
		dest    = fmt.Sprintf("%s/%s.go", mockDir, strcase.ToSnake(decl.SourceName))
	)
	return &Target{
		SrcPkg:  srcPkg,
		SrcName: decl.SourceName,
		MockPkg: mockPkg,
		MockDir: mockDir,
		Dest:    dest,
	}
}
