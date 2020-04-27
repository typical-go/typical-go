package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Context of typical build tool
type Context struct {
	*Descriptor

	TypicalTmp string
	ProjectPkg string

	AppDirs    []string
	AppFiles   []string
	AppSources []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (c *Context, err error) {
	if d == nil {
		return nil, errors.New("TypicalContext: Descriptor can't be empty")
	}
	if err := d.Validate(); err != nil {
		return nil, err
	}

	appSources := d.AppSources()
	if err = validateSources(appSources); err != nil {
		return
	}

	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		appSources = append(appSources, "pkg")
	}

	var appDirs, appFiles []string
	for _, dir := range appSources {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info != nil && info.IsDir() {
				appDirs = append(appDirs, path)
				return nil
			}
			if isWalkTarget(path) {
				appFiles = append(appFiles, path)
			}
			return nil
		})
	}

	return &Context{
		Descriptor: d,
		TypicalTmp: DefaultTypicalTmp,
		ProjectPkg: DefaultProjectPkg,
		AppSources: appSources,
		AppDirs:    appDirs,
		AppFiles:   appFiles,
	}, nil
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}

func validateSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return fmt.Errorf("Source '%s' is not exist", source)
		}
	}
	return
}
