package typmock

import "github.com/typical-go/typical-go/pkg/typannot"

// Mockery is art of mocking
type Mockery struct {
	targetMap  map[string][]*typannot.Mock
	projectPkg string
}

// NewMockery return new instance of store
func NewMockery(projectPkg string) *Mockery {
	return &Mockery{
		targetMap:  make(map[string][]*typannot.Mock),
		projectPkg: projectPkg,
	}
}

// Put target to mockery
func (b *Mockery) Put(target *typannot.Mock) {
	key := target.Pkg
	if _, ok := b.targetMap[key]; ok {
		b.targetMap[key] = append(b.targetMap[key], target)
	} else {
		b.targetMap[key] = []*typannot.Mock{target}
	}
}

// TargetMap contain package and target to be mock
func (b *Mockery) TargetMap(pkgs ...string) map[string][]*typannot.Mock {
	if len(pkgs) > 0 {
		targetMap := make(map[string][]*typannot.Mock)
		for _, pkg := range pkgs {
			if _, ok := b.targetMap[pkg]; ok {
				targetMap[pkg] = b.targetMap[pkg]
			}
		}
		return targetMap
	}
	return b.targetMap
}
