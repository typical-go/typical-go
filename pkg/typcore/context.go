package typcore

import (
	"os"
	"strings"
)

// TypicalContext is context of typical build tool
type TypicalContext struct {
	*Descriptor
	ProjectLayout
	Dirs  []string
	Files []string
}

// DefaultLayout is default project layout
var DefaultLayout = ProjectLayout{
	Cmd:     "cmd",
	Bin:     "bin",
	Temp:    ".typical-tmp",
	Mock:    "mock",
	Release: "release",
}

// ProjectLayout is reflect folder structure of the project
type ProjectLayout struct {
	Bin     string
	Cmd     string
	Temp    string // TODO: temp folder is not part project layout as it is constant for all typical-go
	Mock    string // TODO: mock folder is not part project layout but rather mock generator
	Release string // TODO: consider release folder as project layout
}

// Walk function
func (b *TypicalContext) addFile(path string, info os.FileInfo, err error) error {
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
