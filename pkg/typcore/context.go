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
	BinFolder  string
	CmdFolder  string
	TempFolder string

	Dirs           []string
	Files          []string
	ModulePackage  string
	ProjectSources []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (*Context, error) {
	c := &Context{
		Descriptor: d,

		CmdFolder:  DefaultCmdFolder,
		BinFolder:  DefaultBinFolder,
		TempFolder: DefaultTempFolder,

		ModulePackage:  DefaultModulePackage,
		ProjectSources: RetrieveProjectSources(d),
	}
	for _, dir := range c.ProjectSources {
		if err := filepath.Walk(dir, c.addFile); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Validate typical context
func (c *Context) Validate() error {
	if c.Descriptor == nil {
		return errors.New("TypicalContext: Descriptor can't be empty")
	}
	if err := c.Descriptor.Validate(); err != nil {
		return err
	}

	if c.ModulePackage == "" {
		return errors.New("TypicalContext: ModulePackage can't be empty")
	}

	if err := validateProjectSources(c.ProjectSources); err != nil {
		return fmt.Errorf("TypicalContext: %w", err)
	}
	return nil
}

// TypicalPackage return package of typical
func (c *Context) TypicalPackage() string {
	return fmt.Sprintf("%s/typical", c.ModulePackage)
}

func (c *Context) addFile(path string, info os.FileInfo, err error) error {
	if info != nil {
		if info.IsDir() {
			c.Dirs = append(c.Dirs, path)
			return nil
		}
	}

	if isWalkTarget(path) {
		c.Files = append(c.Files, path)
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
