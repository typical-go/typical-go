package typcore

import (
	"os"
	"strings"
)

// Build is interface of build
type Build interface {
	Run(*BuildContext) error
}

// BuildContext is context of prebuild
type BuildContext struct {
	*Descriptor
	ProjectLayout
	Dirs  []string
	Files []string
}

// Walk function
func (b *BuildContext) addFile(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		b.Dirs = append(b.Dirs, path)
	} else if isWalkTarget(path) {
		b.Files = append(b.Files, path)
	}
	return nil
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
