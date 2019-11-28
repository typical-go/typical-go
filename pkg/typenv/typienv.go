package typenv

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

var (
	appVar        = EnvVar{"TYPICAL_APP", "app"}
	buildToolVar  = EnvVar{"TYPICAL_BUILD_TOOL", "build-tool"}
	prebuilderVar = EnvVar{"TYPICAL_PREBUILDER", "pre-builder"}
	binVar        = EnvVar{"TYPICAL_BIN", "bin"}
	cmdVar        = EnvVar{"TYPICAL_CMD", "cmd"}
	mockVar       = EnvVar{"TYPICAL_MOCK", "mock"}
	releaseVar    = EnvVar{"TYPICAL_RELEASE", "release"}
	dependencyVar = EnvVar{"TYPICAL_DEPENDENCY", "dependency"}
	metadataVar   = EnvVar{"TYPICAL_METADATA", ".typical-metadata"}
)

var (
	Layout struct {
		App      string
		Bin      string
		Cmd      string
		Metadata string
		Mock     string
		Release  string
	}
	Readme string

	BuildTool        string
	BuildToolBin     string
	BuildToolMainPkg string

	Prebuilder        string
	PrebuilderBin     string
	PrebuilderMainPkg string

	Dependency    string
	DependencyPkg string
)

func init() {
	Layout.App = appVar.Value()
	Layout.Cmd = cmdVar.Value()
	Layout.Bin = binVar.Value()
	Layout.Mock = mockVar.Value()
	Layout.Release = releaseVar.Value()
	Layout.Metadata = metadataVar.Value()
	Readme = "README.md"

	BuildTool = buildToolVar.Value()
	BuildToolMainPkg = Layout.Cmd + "/" + BuildTool
	BuildToolBin = Layout.Bin + "/" + BuildTool

	Prebuilder = prebuilderVar.Value()
	PrebuilderMainPkg = Layout.Cmd + "/" + Prebuilder
	PrebuilderBin = Layout.Bin + "/" + Prebuilder

	Dependency = dependencyVar.Value()
	DependencyPkg = "internal/" + Dependency
}

// AppMainPkg return main package of application
func AppMainPkg(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Cmd, strcase.ToKebab(name))
}

// AppBin return bin path of application
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", Layout.Bin, strcase.ToKebab(name))
}
