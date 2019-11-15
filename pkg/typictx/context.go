package typictx

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/utility/collection"
)

// Context of typical application
type Context struct {
	Name            string
	Description     string
	Package         string
	AppModule       AppModule
	Modules         collection.Interfaces
	Release         Release
	TestTargets     collection.Strings
	MockTargets     collection.Strings
	Constructors    collection.Interfaces
	ReadmeGenerator interface {
		GenerateReadme(*Context) (err error)
	}
}

// AppModule is application module
type AppModule interface {
	Run() interface{}
}

// Validate context
func (c *Context) Validate() (err error) {
	if c.Name == "" {
		return invalidContextError("Name can't not empty")
	}
	if c.Package == "" {
		return invalidContextError("Package can't not empty")
	}
	if err = c.Release.Validate(); err != nil {
		return fmt.Errorf("Release: %s", err.Error())
	}
	return nil
}

// AllModule return app module and modules
func (c *Context) AllModule() (modules []interface{}) {
	modules = append(modules, c.AppModule)
	modules = append(modules, c.Modules...)
	return
}
