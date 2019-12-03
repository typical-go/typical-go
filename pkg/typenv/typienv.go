package typenv

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

var (
// appVar        = EnvVar{"TYPICAL_APP", "app"}
// buildToolVar  = EnvVar{"TYPICAL_BUILD_TOOL", "build-tool"}
// prebuilderVar = EnvVar{"TYPICAL_PREBUILDER", "pre-builder"}
// binVar        = EnvVar{"TYPICAL_BIN", "bin"}
// cmdVar        = EnvVar{"TYPICAL_CMD", "cmd"}
// mockVar       = EnvVar{"TYPICAL_MOCK", "mock"}
// releaseVar    = EnvVar{"TYPICAL_RELEASE", "release"}
// dependencyVar = EnvVar{"TYPICAL_DEPENDENCY", "dependency"}
// metadataVar   = EnvVar{"TYPICAL_METADATA", ".typical-metadata"}
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

	BuildTool        = "build-tool"
	BuildToolBin     = Layout.Bin + "/" + BuildTool
	BuildToolMainPkg = Layout.Cmd + "/" + BuildTool

	Prebuilder        = "pre-builder"
	PrebuilderBin     = Layout.Cmd + "/" + Prebuilder
	PrebuilderMainPkg = Layout.Bin + "/" + Prebuilder

	Dependency    = "dependency"
	DependencyPkg = "internal/" + Dependency
)

// AppMainPkg return main package of application
func AppMainPkg(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Cmd, strcase.ToKebab(name))
}

// AppBin return bin path of application
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Bin, strcase.ToKebab(name))
}
