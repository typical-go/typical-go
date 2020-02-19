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
	Package     string
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

// Configuration is interface of configuration
type Configuration interface {
	Store() *ConfigStore
	Setup() error
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

	var root string
	if root, err = os.Getwd(); err != nil {
		return errors.New("Descriptor: Fail to get working directory")
	}

	if d.Name == "" {
		d.Name = filepath.Base(root)
	} else if err = validateName(d.Name); err != nil {
		return
	}

	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if len(d.Sources) < 1 {
		// TODO: sources provided by app type or the package name where app belongs
		d.Sources = []string{"app", "pkg"}
	}

	if err = validateSources(d.Sources); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	// TODO: then validate if sources if exist in the project

	if d.Package == "" {
		// TODO: get package from gomod file
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

func validateName(name string) (err error) {
	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return errors.New("Descriptor: Invalid `Name`")
	}
	return
}

func validateSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return
		}
	}
	return
}
