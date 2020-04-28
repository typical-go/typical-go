package typmock

import (
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typast"
)

// Mockery is art of mocking
type Mockery struct {
	targetMap  map[string][]*Target
	projectPkg string
}

// Target for mockery
type Target struct {
	Dir    string
	Pkg    string
	Source string
	Parent string
}

// NewMockery return new instance of store
func NewMockery(projectPkg string) *Mockery {
	return &Mockery{
		targetMap:  make(map[string][]*Target),
		projectPkg: projectPkg,
	}
}

// CreateMockery to get mock annotated
func CreateMockery(ast *typast.ASTStore, projectPkg string) (b *Mockery, err error) {
	b = NewMockery(projectPkg)
	for _, a := range ast.Annots {
		if a.Equal("mock", typast.InterfaceType) {
			pkg := a.File.Name.Name
			dir := filepath.Dir(a.Path)

			b.Put(&Target{
				Dir:    dir,
				Pkg:    pkg,
				Source: a.SourceName,
				Parent: dir[:len(dir)-len(pkg)],
			})
		}
	}

	return
}

// Put target to mockery
func (b *Mockery) Put(target *Target) {
	key := target.Pkg
	if _, ok := b.targetMap[key]; ok {
		b.targetMap[key] = append(b.targetMap[key], target)
	} else {
		b.targetMap[key] = []*Target{target}
	}
}

// TargetMap contain package and target to be mock
func (b *Mockery) TargetMap(pkgs ...string) map[string][]*Target {
	if len(pkgs) > 0 {
		targetMap := make(map[string][]*Target)
		for _, pkg := range pkgs {
			if _, ok := b.targetMap[pkg]; ok {
				targetMap[pkg] = b.targetMap[pkg]
			}
		}
		return targetMap
	}
	return b.targetMap
}
