package typbuildtool

import "github.com/typical-go/typical-go/pkg/buildkit"

// MockStore responsible to store mock target
type MockStore struct {
	m map[string][]*buildkit.GoMock
}

// NewMockStore to return new instance of MockStore
func NewMockStore() *MockStore {
	return &MockStore{
		m: make(map[string][]*buildkit.GoMock),
	}
}

// Put new target
func (b *MockStore) Put(target *buildkit.GoMock) {
	key := target.MockDir
	if _, ok := b.m[key]; ok {
		b.m[key] = append(b.m[key], target)
	} else {
		b.m[key] = []*buildkit.GoMock{target}
	}
}

// Map of mock-target
func (b *MockStore) Map() map[string][]*buildkit.GoMock {
	return b.m
}
