package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

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

	return
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
