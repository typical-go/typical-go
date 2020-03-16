package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
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

// RetrieveProjectSources to retrieve project source
func RetrieveProjectSources(d *Descriptor) (sources []string, err error) {
	if sourceable, ok := d.App.(SourceableApp); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, RetrievePackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}
	if err = validateProjectSources(sources); err != nil {
		return nil, err
	}
	return
}

func validateProjectSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return fmt.Errorf("Source '%s' is not exist", source)
		}
	}
	return
}

// RetrievePackageName return package name of the interface
func RetrievePackageName(v interface{}) string {
	if common.IsNil(v) {
		return ""
	}
	s := reflect.TypeOf(v).String()
	if dot := strings.Index(s, "."); dot > 0 {
		if strings.HasPrefix(s, "*") {
			return s[1:dot]
		}
		return s[:dot]
	}
	return ""
}
