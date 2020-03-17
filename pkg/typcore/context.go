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

	ProjectPackage string
	ProjectDirs    []string
	ProjectFiles   []string
	ProjectSources []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (c *Context, err error) {
	if d == nil {
		return nil, errors.New("TypicalContext: Descriptor can't be empty")
	}
	if err := d.Validate(); err != nil {
		return nil, err
	}

	var projectSources []string
	if projectSources, err = RetrieveProjectSources(d); err != nil {
		return nil, err
	}

	var projectDirs, projectFiles []string
	for _, dir := range projectSources {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info != nil && info.IsDir() {
				projectDirs = append(projectDirs, path)
				return nil
			}
			if isWalkTarget(path) {
				projectFiles = append(projectFiles, path)
			}
			return nil
		})
	}

	return &Context{
		Descriptor:     d,
		ProjectPackage: DefaultProjectPackage,
		ProjectSources: projectSources,
		ProjectDirs:    projectDirs,
		ProjectFiles:   projectFiles,
	}, nil
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
