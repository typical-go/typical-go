package typcore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TypicalContext is context of typical build tool
type TypicalContext struct {
	*Descriptor
	BinFolder  string
	CmdFolder  string
	TempFolder string // TODO: temp folder is not part project layout as it is constant for all typical-go

	Dirs           []string
	Files          []string
	ModulePackage  string
	ProjectSources []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (*TypicalContext, error) {
	c := &TypicalContext{
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
func (t *TypicalContext) Validate() error {
	if t.Descriptor == nil {
		return errors.New("TypicalContext: Descriptor can't be empty")
	}
	if err := t.Descriptor.Validate(); err != nil {
		return err
	}

	if t.ModulePackage == "" {
		return errors.New("TypicalContext: ModulePackage can't be empty")
	}

	if err := validateProjectSources(t.ProjectSources); err != nil {
		return fmt.Errorf("TypicalContext: %w", err)
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

func (t *TypicalContext) addFile(path string, info os.FileInfo, err error) error {
	if info != nil {
		if info.IsDir() {
			t.Dirs = append(t.Dirs, path)
			return nil
		}
	}

	if isWalkTarget(path) {
		t.Files = append(t.Files, path)
	}
	return nil
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
