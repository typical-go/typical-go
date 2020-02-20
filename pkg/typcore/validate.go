package typcore

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
)

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

	if d.ModulePackage == "" {
		if d.ModulePackage, err = defaultModulePackage(root); err != nil {
			return
		}
	}

	if d.App == nil {
		return errors.New("Descriptor: App can't be nil")
	} else if err = common.Validate(d.App); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	if d.BuildTool == nil {
		return errors.New("Descriptor: BuildTool can't be nil")
	} else if err = common.Validate(d.BuildTool); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	if len(d.ProjectSources) < 1 {
		d.ProjectSources = d.defaultProjectSources()
	}
	if err = validateProjectSources(d.ProjectSources); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	return
}

func (d *Descriptor) defaultProjectSources() (sources []string) {
	if sourceable, ok := d.App.(Sourceable); ok {
		sources = append(sources, sourceable.Sources()...)
	} else {
		sources = append(sources, pkgName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}
	return
}

func defaultModulePackage(root string) (pkg string, err error) {
	var (
		gomod *common.GoMod
	)
	if gomod, err = common.CreateGoMod(root + "/go.mod"); err != nil {
		// NOTE: go.mod is not exist. Check if the project sit in $GOPATH
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			return root[len(gopath):], nil
		}

		return "", errors.New("`go.mod` is missing and the project not in $GOPATH")
	}

	return gomod.ModulePackage, nil
}

func validateName(name string) (err error) {
	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return errors.New("Descriptor: Invalid `Name`")
	}
	return
}

func validateProjectSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return fmt.Errorf("Source '%s' is not exist", source)
		}
	}
	return
}

func pkgName(v interface{}) string {
	s := reflect.TypeOf(v).String()
	if dot := strings.Index(s, "."); dot > 0 {
		if strings.HasPrefix(s, "*") {
			return s[1:dot]
		}
		return s[:dot]
	}
	return ""
}
