package typast

import (
	"os"
	"path/filepath"
	"strings"
)

type (
	Walker interface {
		Walk() []string
	}
	Layouts   []string
	FilePaths []string
)

//
// Layouts
//

var _ Walker = (Layouts)(nil)

func (l Layouts) Walk() []string {
	var filePaths []string
	for _, layout := range l {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	}
	return filePaths
}

//
// FilePaths
//

var _ Walker = (FilePaths)(nil)

func (f FilePaths) Walk() []string {
	return f
}
