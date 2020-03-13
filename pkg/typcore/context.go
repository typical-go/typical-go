package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typast"
)

// Context of typical build tool
type Context struct {
	*Descriptor
	BinFolder  string
	CmdFolder  string
	TempFolder string

	ProjectDirs    []string
	ProjectFiles   []string
	ProjectPackage string
	ProjectSources []string

	ast *typast.Ast
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (*Context, error) {
	c := &Context{
		Descriptor: d,

		CmdFolder:  DefaultCmdFolder,
		BinFolder:  DefaultBinFolder,
		TempFolder: DefaultTempFolder,

		ProjectPackage: DefaultProjectPackage,
		ProjectSources: RetrieveProjectSources(d),
	}
	for _, dir := range c.ProjectSources {
		if err := filepath.Walk(dir, c.addFile); err != nil {
			return nil, err
		}
	}
	return c, nil
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
