package filekit

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v2"
)

// MatchMulti similar with filepath.Match except accept multiple pattern
func MatchMulti(patterns []string, name string) bool {
	for _, pattern := range patterns {
		if matched, _ := doublestar.Match(pattern, name); matched {
			return true
		}
	}
	return false
}

// FindDir find directory
func FindDir(includes, excludes []string) (packages []string, err error) {
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if MatchMulti(includes, path) && !MatchMulti(excludes, path) && info.IsDir() {
			packages = append(packages, "./"+path)
		}
		return nil
	})
	return
}
