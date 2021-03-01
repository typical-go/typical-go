package envkit

import "sort"

type (
	// Map contain environment variable
	Map map[string]string
)

// SortedKeys of EnvMap
func (m Map) SortedKeys() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
