package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"
)

// Context of typical build tool
type Context struct {
	*Descriptor
	TempFolder string

	ProjectDirs    []string
	ProjectFiles   []string
	ProjectPackage string
	ProjectSources []string

	ast *typast.Ast

	Logger
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) *Context {
	c := &Context{
		Descriptor: d,

		TempFolder: DefaultTempFolder,

		ProjectPackage: DefaultProjectPackage,
		ProjectSources: RetrieveProjectSources(d),
		Logger:         NewLogger(),
	}
	for _, dir := range c.ProjectSources {
		filepath.Walk(dir, c.addFile)
	}
	return c
}

// Ast contain detail of AST analysis
func (c *Context) Ast() *typast.Ast {
	if c.ast == nil {
		var err error
		if c.ast, err = typast.Walk(c.ProjectFiles); err != nil {
			log.Errorf("PreconditionContext: %w", err.Error())
		}
	}
	return c.ast
}

// Validate typical context
func (c *Context) Validate() error {
	if c.Descriptor == nil {
		return errors.New("TypicalContext: Descriptor can't be empty")
	}
	if err := c.Descriptor.Validate(); err != nil {
		return err
	}

	if c.ProjectPackage == "" {
		return errors.New("TypicalContext: ProjectPackage can't be empty")
	}

	if err := validateProjectSources(c.ProjectSources); err != nil {
		return fmt.Errorf("TypicalContext: %w", err)
	}
	return nil
}

func (c *Context) addFile(path string, info os.FileInfo, err error) error {
	if info != nil {
		if info.IsDir() {
			c.ProjectDirs = append(c.ProjectDirs, path)
			return nil
		}
	}

	if isWalkTarget(path) {
		c.ProjectFiles = append(c.ProjectFiles, path)
	}
	return nil
}

func validateProjectSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return fmt.Errorf("Source '%s' is not exist", source)
		}
	}
	return
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}

// RetrieveProjectSources to retrieve project source
func RetrieveProjectSources(d *Descriptor) (sources []string) {
	if sourceable, ok := d.App.(SourceableApp); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, RetrievePackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
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
