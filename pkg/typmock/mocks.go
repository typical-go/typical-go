package typmock

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

const DefaultParent = "internal/generated"

type (
	// Mocks is art of mocking
	Mocks []*Mock
	// Mock annotation data
	Mock struct {
		Pkg     string `json:"-"`
		Source  string `json:"-"`
		Dest    string `json:"-"`
		MockPkg string `json:"-"`
	}
)

// MockGen execute mockgen bash
func MockGen(c *typgo.Context, destPkg, dest, srcPkg, src string) error {
	mockgen, err := typgo.InstallTool(c, "mockgen", "github.com/golang/mock/mockgen")
	if err != nil {
		return err
	}

	return c.ExecuteCommand(&typgo.Command{
		Name: mockgen,
		Args: []string{
			"-destination", dest,
			"-package", destPkg,
			srcPkg,
			src,
		},
		Stderr: os.Stderr,
	})
}

// CreateMock to create mock
func CreateMock(d *typgen.Annotation) *Mock {
	source := d.Decl.GetName()

	return &Mock{
		Pkg:     d.PackagePath(),
		Source:  source,
		Dest:    fmt.Sprintf("%s/%s.go", GeneratedDir(d, "mock"), strcase.ToSnake(source)),
		MockPkg: d.Package() + "_mock",
	}
}

func GeneratedDir(d *typgen.Annotation, suffix string) string {
	dir := filepath.Dir(d.Path())
	if dir == "." {
		return DefaultParent
	}
	dir = strings.ReplaceAll(dir, "internal/", "")
	return DefaultParent + "/" + dir + "_" + suffix
}
