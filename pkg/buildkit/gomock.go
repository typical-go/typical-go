package buildkit

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typast"
)

// GoMock to be mocked
type GoMock struct {
	SrcDir  string
	SrcPkg  string
	SrcName string
	MockPkg string
	MockDir string
	Dest    string

	tmpFolder string
}

// CreateGoMock to create new instance of GoMock
func CreateGoMock(tmpFolder, projectPackage string, decl *typast.Declaration) *GoMock {
	pkg := decl.File.Name.Name
	dir := filepath.Dir(decl.Path)
	srcName := decl.SourceName

	dirDest := dir[:len(dir)-len(pkg)]
	srcPkg := fmt.Sprintf("%s/%s", projectPackage, dir)
	mockPkg := fmt.Sprintf("mock_%s", pkg)
	mockDir := fmt.Sprintf("%s%s", dirDest, mockPkg)
	dest := fmt.Sprintf("%s/%s.go", mockDir, strcase.ToSnake(srcName))

	return &GoMock{
		tmpFolder: tmpFolder,
		SrcPkg:    srcPkg,
		SrcName:   srcName,
		MockPkg:   mockPkg,
		MockDir:   mockDir,
		Dest:      dest,
	}

}

// Execute gomock
func (m *GoMock) Execute(ctx context.Context) (err error) {
	mockgen := fmt.Sprintf("%s/bin/mockgen", m.tmpFolder)

	if _, err = os.Stat(mockgen); os.IsNotExist(err) {
		if err = NewGoBuild(mockgen, "github.com/golang/mock/mockgen").Execute(ctx); err != nil {
			return
		}
	}

	cmd := exec.CommandContext(ctx, mockgen,
		"-destination", m.Dest,
		"-package", m.MockPkg,
		m.SrcPkg,
		m.SrcName,
	)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (m *GoMock) String() string {
	return fmt.Sprintf("%s.%s", m.SrcPkg, m.SrcName)
}
