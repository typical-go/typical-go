package typcore

import (
	"os"
	"path/filepath"
	"strings"
)

// ProjectInfo is detail of project
type ProjectInfo struct {
	Dirs, Files []string
}

// AppendDir to append dir
func (p *ProjectInfo) AppendDir(dirs ...string) *ProjectInfo {
	p.Dirs = append(p.Dirs, dirs...)
	return p
}

// AppendFile to append file
func (p *ProjectInfo) AppendFile(files ...string) *ProjectInfo {
	p.Files = append(p.Files, files...)
	return p
}

// Walk function
func (p *ProjectInfo) Walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		p.AppendDir(path)
	} else if isWalkTarget(path) {
		p.AppendFile(path)
	}
	return nil
}

// ReadProject to read the project to get Project Info
func ReadProject(root string) (proj ProjectInfo, err error) {
	proj.AppendDir(root)
	err = filepath.Walk(root, proj.Walk)
	return
}

func isWalkTarget(filename string) bool { //  TODO: move out from walker package
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
