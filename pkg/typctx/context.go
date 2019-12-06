package typctx

import (
	"fmt"
	"io"

	"github.com/typical-go/typical-go/pkg/typcfg"

	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Context of typical application
type Context struct {
	Name         string
	Description  string
	Package      string
	Version      string
	AppModule    interface{}
	Modules      coll.Interfaces
	ConfigLoader typcfg.Loader
	Releaser     *typrls.Releaser

	TestTargets  coll.Strings
	MockTargets  coll.Strings
	Constructors coll.Interfaces

	ReadmeGenerator interface {
		Generate(*Context, io.Writer) error
	}
}

// Validate context
func (c *Context) Validate() (err error) {
	if c.Name == "" {
		return invalidContextError("Name can't be empty")
	}
	if c.Version == "" {
		c.Version = "0.0.1"
	}
	if c.Package == "" {
		return invalidContextError("Package can't be empty")
	}
	if c.Releaser != nil {
		if err = c.Releaser.Validate(); err != nil {
			return fmt.Errorf("Releaser: %s", err.Error())
		}
	}
	if c.ConfigLoader == nil {
		c.ConfigLoader = typcfg.DefaultLoader()
	}
	return
}

// AllModule return app module and modules
func (c *Context) AllModule() (modules []interface{}) {
	// NOTE: modules should be before app module to make sure it readiness
	modules = append(modules, c.Modules...)
	modules = append(modules, c.AppModule)
	return
}
