package typbuildtool

// MockStore responsible to store mock target
type MockStore struct {
	m map[string][]*MockTarget
}

// MockTarget to be mocked
type MockTarget struct {
	SrcDir  string
	SrcPkg  string
	SrcName string
	MockPkg string
	MockDir string
	Dest    string
}

// NewMockStore to return new instance of MockStore
func NewMockStore() *MockStore {
	return &MockStore{
		m: make(map[string][]*MockTarget),
	}
}

// Put new target
func (b *MockStore) Put(target *MockTarget) {
	key := target.MockDir
	if _, ok := b.m[key]; ok {
		b.m[key] = append(b.m[key], target)
	} else {
		b.m[key] = []*MockTarget{target}
	}
}

// Map of mock-target
func (b *MockStore) Map() map[string][]*MockTarget {
	return b.m
}
