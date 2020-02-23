package typcore

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
)

var (
	// DefaultModulePackage is default value for ModulePackage
	DefaultModulePackage = "" // NOTE: supply by ldflags

	// DefaultTempFolder is default value for temp folder location
	DefaultTempFolder = ".typical-tmp"

	// DefaultCmdFolder is default value for cmd folder location
	DefaultCmdFolder = "cmd"

	// DefaultBinFolder is default value for bin folder location
	DefaultBinFolder = "bin"

	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"
)

// TypicalContext is context of typical build tool
type TypicalContext struct {
	*Descriptor
	BinFolder     string
	CmdFolder     string
	TempFolder    string // TODO: temp folder is not part project layout as it is constant for all typical-go
	MockFolder    string // TODO: mock folder is not part project layout but rather mock generator
	ReleaseFolder string // TODO: consider release folder as project layout

	Dirs           []string
	Files          []string
	ModulePackage  string
	ProjectSources []string
}

// CreateContext return new constructor of TypicalContext
func CreateContext(d *Descriptor) (*TypicalContext, error) {
	projectSources := defaultProjectSources(d)
	c := &TypicalContext{
		Descriptor: d,

		CmdFolder:     DefaultCmdFolder,
		BinFolder:     DefaultBinFolder,
		TempFolder:    DefaultTempFolder,
		MockFolder:    "mock",
		ReleaseFolder: "release",

		ModulePackage:  DefaultModulePackage,
		Dirs:           projectSources,
		ProjectSources: projectSources,
	}
	for _, dir := range c.Dirs {
		if err := filepath.Walk(dir, c.addFile); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Validate typical context
func (t *TypicalContext) Validate() (err error) {
	if t.ModulePackage == "" {
		return errors.New("TypicalContext: ModulePackage can't be empty")
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

func defaultProjectSources(d *Descriptor) (sources []string) {
	if sourceable, ok := d.App.(Sourceable); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, common.PackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}
	return
}
