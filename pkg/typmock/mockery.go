package typmock

import (
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typast"
)

type (
	// Mockery is art of mocking
	Mockery struct {
		Map        Map    `json:"target_map"`
		ProjectPkg string `json:"project_pkg"`
	}
	// Map of mock target
	Map map[string][]*Mock
	// Mock annotation data
	Mock struct {
		Dir     string `json:"-"`
		Pkg     string `json:"-"`
		Source  string `json:"-"`
		Parent  string `json:"-"`
		MockPkg string `json:"-"`
	}
)

// NewMockery return new instancen of Mockery
func NewMockery(projectPkg string) *Mockery {
	return &Mockery{
		Map:        make(map[string][]*Mock),
		ProjectPkg: projectPkg,
	}
}

// Put target to mockery
func (m Mockery) Put(target *Mock) {
	key := target.Dir
	if _, ok := m.Map[key]; ok {
		m.Map[key] = append(m.Map[key], target)
	} else {
		m.Map[key] = []*Mock{target}
	}
}

// Filter contain package and target to be mock
func (m Mockery) Filter(pkgs ...string) Map {
	targetMap := make(map[string][]*Mock)
	for _, pkg := range pkgs {
		if _, ok := m.Map[pkg]; ok {
			targetMap[pkg] = m.Map[pkg]
		}
	}
	return targetMap
}

// CreateMock to create mock
func CreateMock(annot *typast.Directive) *Mock {
	pkg := annot.Decl.Package
	dir := filepath.Dir(annot.Decl.Path)

	parent := ""
	if dir != "." {
		parent = dir[:len(dir)-len(pkg)]
	}

	return &Mock{
		Dir:     dir,
		Pkg:     pkg,
		Source:  annot.GetName(),
		Parent:  parent,
		MockPkg: fmt.Sprintf("%s_mock", pkg),
	}
}
