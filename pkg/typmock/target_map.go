package typmock

// TargetMap of mock
type TargetMap map[string][]*Mock

// Put target to mockery
func (m TargetMap) Put(target *Mock) {
	key := target.Dir
	if _, ok := m[key]; ok {
		m[key] = append(m[key], target)
	} else {
		m[key] = []*Mock{target}
	}
}

// Filter contain package and target to be mock
func (m TargetMap) Filter(pkgs ...string) TargetMap {
	targetMap := make(map[string][]*Mock)
	for _, pkg := range pkgs {
		if _, ok := m[pkg]; ok {
			targetMap[pkg] = m[pkg]
		}
	}
	return targetMap
}
