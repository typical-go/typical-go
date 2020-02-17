package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
)

// Descriptor describe the project
type Descriptor struct {
	Name        string
	Description string
	Package     string // TODO: get package from gomod file
	Version     string
	Sources     []string

	App           App
	BuildTool     BuildTool
	Configuration Configuration
}

// App is interface of app
type App interface {
	Run(*Descriptor) error
}

// BuildTool interface
type BuildTool interface {
	Run(*TypicalContext) error
}

// RunApp to run app
func (d *Descriptor) RunApp() (err error) {
	if d.App == nil {
		return errors.New("Descriptor is missing `App`")
	}
	if err = d.Validate(); err != nil {
		return
	}
	return d.App.Run(d)
}

// RunBuild to run build
func (d *Descriptor) RunBuild() (err error) {
	if d.BuildTool == nil {
		return errors.New("Descriptor is missing `BuildTool`")
	}
	if err = d.Validate(); err != nil {
		return
	}
	c := &TypicalContext{
		Descriptor:    d,
		ProjectLayout: DefaultLayout,
		Dirs:          d.Sources,
	}
	for _, dir := range c.Dirs {
		if err = filepath.Walk(dir, c.addFile); err != nil {
			return
		}
	}
	if d.Configuration != nil {
		if err = d.Configuration.Setup(); err != nil {
			return
		}
	}
	return d.BuildTool.Run(c)
}

// Validate context
func (d *Descriptor) Validate() (err error) {

	if d.Name == "" {
		d.Name = defaultName()
	} else {
		r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
		if !r.MatchString(d.Name) {
			return errors.New("Descriptor: Invalid `Name`")
		}
	}

	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if len(d.Sources) < 1 {
		// TODO: sources provided by app type or the package name where app belongs
		d.Sources = []string{"app", "pkg"}

		// TODO: then validate if sources if exist in the project
	}

	if d.Package == "" {
		return errors.New("Descriptor: Package can't be empty")
	}

	if err = common.Validate(d.BuildTool); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	if err = common.Validate(d.App); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	return
}

func defaultName() (s string) {
	var err error
	if s, err = os.Getwd(); err != nil {
		return "noname"
	}
	return filepath.Base(s)
}
