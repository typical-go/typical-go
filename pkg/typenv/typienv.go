package typenv

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

var (
	Layout = struct {
		App      string
		Bin      string
		Cmd      string
		Metadata string
		Mock     string
		Release  string
	}{
		App:      "app",
		Cmd:      "cmd",
		Bin:      "bin",
		Metadata: ".typical-metadata",
		Mock:     "mock",
		Release:  "release",
	}
	Readme = "README.md"

	BuildToolBin     = Layout.Bin + "/build-tool"
	BuildToolMainPkg = Layout.Cmd + "/build-tool"

	PrebuilderBin     = Layout.Cmd + "/pre-builder"
	PrebuilderMainPkg = Layout.Bin + "/pre-builder"

	DependencyPkg = "internal/dependency"
)

// AppMainPkg return main package of application
func AppMainPkg(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Cmd, strcase.ToKebab(name))
}

// AppBin return bin path of application
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Bin, strcase.ToKebab(name))
}
