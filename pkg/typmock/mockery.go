package typmock

var (
	// MockTag is tag for mock
	MockTag = "mock"
)

type (
	// Mockery is art of mocking
	Mockery struct {
		TargetMap  TargetMap `json:"target_map"`
		ProjectPkg string    `json:"project_pkg"`
	}

	// TargetMap of mock
	TargetMap map[string][]*Mock

	// Mock annotation data
	Mock struct {
		Dir     string `json:"-"`
		Pkg     string `json:"-"`
		Source  string `json:"-"`
		Parent  string `json:"-"`
		MockPkg string `json:"-"`
	}
)

// Put target to mockery
func (b *Mockery) Put(target *Mock) {
	key := target.Dir
	if _, ok := b.TargetMap[key]; ok {
		b.TargetMap[key] = append(b.TargetMap[key], target)
	} else {
		b.TargetMap[key] = []*Mock{target}
	}
}

// Filter contain package and target to be mock
func (b *Mockery) Filter(pkgs ...string) TargetMap {
	targetMap := make(map[string][]*Mock)
	for _, pkg := range pkgs {
		if _, ok := b.TargetMap[pkg]; ok {
			targetMap[pkg] = b.TargetMap[pkg]
		}
	}
	return targetMap

}
