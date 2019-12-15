package typcore

import (
	"errors"
	"fmt"
	"io"
	"reflect"

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
	ConfigLoader Loader
	Releaser     *typrls.Releaser

	MockTargets  coll.Strings
	Constructors coll.Interfaces

	ReadmeGenerator interface {
		Generate(*Context, io.Writer) error
	}
}

// Validate context
func (c *Context) Validate() (err error) {
	if c.Name == "" {
		return errors.New("Context: Name can't be empty")
	}
	if c.Package == "" {
		return errors.New("Context: Package can't be empty")
	}
	if c.Version == "" {
		c.Version = "0.0.1"
	}
	if c.ConfigLoader == nil {
		c.ConfigLoader = DefaultLoader()
	}
	if err = validate(c.Releaser); err != nil {
		return fmt.Errorf("Context: Releaser: %w", err)
	}
	for _, module := range c.AllModule() {
		if err = validate(module); err != nil {
			return fmt.Errorf("Context: %s: %w", Name(module), err)
		}
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

func validate(v interface{}) (err error) {
	if isNil(v) {
		return
	}
	if validator, ok := v.(Validator); ok {
		if err = validator.Validate(); err != nil {
			return
		}
	}
	return
}

func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
