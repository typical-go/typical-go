package typcore

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Context of typical build tool
type Context struct {
	*Descriptor

	TypicalTmp string
	ProjectPkg string

	AppDirs  []string
	AppFiles []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (c *Context, err error) {
	if d == nil {
		return nil, errors.New("TypicalContext: Descriptor can't be empty")
	}

	if err := d.Validate(); err != nil {
		return nil, err
	}

	var appDirs, appFiles []string

	for _, layout := range d.Layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				appDirs = append(appDirs, path)
				return nil
			}

			if isGoSource(path) {
				appFiles = append(appFiles, path)
			}
			return nil
		})
	}

	return &Context{
		Descriptor: d,
		TypicalTmp: DefaultTypicalTmp,
		ProjectPkg: DefaultProjectPkg,
		AppDirs:    appDirs,
		AppFiles:   appFiles,
	}, nil
}

func isGoSource(path string) bool {

	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}
