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

const DefaultParent = "internal/generated/mock"

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

	return c.Execute(&typgo.Bash{
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
func CreateMock(annot *typgen.Directive) *Mock {
	dir := filepath.Dir(annot.Decl.Path)
	target := GenTarget(dir)
	mockPkg := fmt.Sprintf("%s_mock", annot.Decl.Package)
	source := annot.GetName()

	return &Mock{
		Pkg:     fmt.Sprintf("%s/%s", typgo.ProjectPkg, dir),
		Source:  source,
		Dest:    fmt.Sprintf("%s/%s/%s.go", target, mockPkg, strcase.ToSnake(source)),
		MockPkg: mockPkg,
	}
}

func GenTarget(dir string) string {
	if dir == "." {
		return DefaultParent
	}
	target := filepath.Dir(dir)
	var words []string
	for _, word := range strings.Split(target, "/") {
		if word != "internal" {
			words = append(words, word)
		}
	}
	words = append(strings.Split(DefaultParent, "/"), words...)
	return strings.Join(words, "/")
}
