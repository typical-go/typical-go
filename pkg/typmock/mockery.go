package typmock

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typvar"
)

var (
	// MockTag is tag for mock
	MockTag = "mock"
)

// Mockery is art of mocking
type Mockery struct {
	TargetMap  TargetMap `json:"target_map"`
	ProjectPkg string    `json:"project_pkg"`
}

func createMockery(c *typgo.Context) *Mockery {
	m := TargetMap{}
	for _, annot := range c.ASTStore.Annots {
		mock := CreateMock(annot)
		if mock != nil {
			m.Put(mock)
		}
	}

	return &Mockery{
		TargetMap:  m,
		ProjectPkg: typvar.ProjectPkg,
	}
}
