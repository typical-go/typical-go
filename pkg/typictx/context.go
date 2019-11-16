package typictx

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/collection"
)

// Context of typical application
type Context struct {
	Name         string
	Description  string
	Package      string
	Version      string
	AppModule    AppModule
	Modules      collection.Interfaces
	TestTargets  collection.Strings
	MockTargets  collection.Strings
	Constructors collection.Interfaces
	Tagging
	ReleaseTargets  []string
	Github          *Github
	ReadmeGenerator interface {
		Generate(*Context, io.Writer) (err error)
	}
}

// Validate context
func (c *Context) Validate() (err error) {
	if c.Name == "" {
		return invalidContextError("Name can't not empty")
	}
	if c.Package == "" {
		return invalidContextError("Package can't not empty")
	}
	if len(c.ReleaseTargets) < 1 {
		return errors.New("Missing 'ReleaseTargets'")
	}
	for _, target := range c.ReleaseTargets {
		if !strings.Contains(target, "/") {
			return fmt.Errorf("Missing '/' in Release Target '%s'", target)
		}
	}
	return nil
}

// AllModule return app module and modules
func (c *Context) AllModule() (modules []interface{}) {
	modules = append(modules, c.AppModule)
	modules = append(modules, c.Modules...)
	return
}
