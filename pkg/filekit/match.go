package filekit

import "path/filepath"

// MatchMulti similar with filepath.Match except accept multiple pattern
func MatchMulti(patterns []string, name string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, name); matched {
			return true
		}
	}
	return false
}
