package typmock

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Mock annotation data
	Mock struct {
		Package            string
		Source             string
		Destination        string
		DestinationPackage string
	}
)

// MockGen execute mockgen bash
func (m *Mock) Generate(c *typgo.Context) error {
	mockgen, err := typgo.InstallTool(c, "mockgen", "github.com/golang/mock/mockgen")
	if err != nil {
		return err
	}

	return c.ExecuteCommand(&typgo.Command{
		Name: mockgen,
		Args: []string{
			"-destination", m.Destination,
			"-package", m.DestinationPackage,
			m.Package,
			m.Source,
		},
		Stderr: os.Stderr,
	})
}

// CreateMock to create mock
func CreateMock(d *typgen.Annotation) *Mock {
	source := d.Decl.GetName()
	dir := typgen.CreateTargetDir(d.Path(), "mock")

	return &Mock{
		Package:            d.PackagePath(),
		Source:             source,
		Destination:        fmt.Sprintf("%s/%s.go", dir, strcase.ToSnake(source)),
		DestinationPackage: d.Package() + "_mock",
	}
}
