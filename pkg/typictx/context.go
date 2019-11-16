package typictx

import (
	"fmt"
	"io"

	"github.com/typical-go/typical-go/pkg/typirelease"
	"github.com/typical-go/typical-go/pkg/utility/collection"
)

// Context of typical application
type Context struct {
	Name            string
	Description     string
	Package         string
	AppModule       AppModule
	Modules         collection.Interfaces
	TestTargets     collection.Strings
	MockTargets     collection.Strings
	Constructors    collection.Interfaces
	ReadmeGenerator interface {
		Generate(*Context, io.Writer) error
	}
	typirelease.Releaser
}

// Validate context
func (c *Context) Validate() (err error) {
	if c.Name == "" {
		return invalidContextError("Name can't not empty")
	}
	if c.Package == "" {
		return invalidContextError("Package can't not empty")
	}
	if err = c.Releaser.Validate(); err != nil {
		return fmt.Errorf("Releaser: %s", err.Error())
	}
	return
}

// AllModule return app module and modules
func (c *Context) AllModule() (modules []interface{}) {
	modules = append(modules, c.AppModule)
	modules = append(modules, c.Modules...)
	return
}
