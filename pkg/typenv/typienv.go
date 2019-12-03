package typenv

import (
	"fmt"
	"os"
	"path/filepath"
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

	App        = name()
	AppBin     = fmt.Sprintf("%s/%s", Layout.Bin, App)
	AppMainPkg = fmt.Sprintf("%s/%s", Layout.Cmd, App)

	BuildToolBin     = fmt.Sprintf("%s/build-tool", Layout.Bin)
	BuildToolMainPkg = fmt.Sprintf("%s/build-tool", Layout.Cmd)

	PrebuilderBin     = fmt.Sprintf("%s/pre-builder", Layout.Bin)
	PrebuilderMainPkg = fmt.Sprintf("%s/pre-builder", Layout.Cmd)

	Dependency    = "dependency"
	DependencyPkg = fmt.Sprintf("internal/%s", Dependency)
)

func name() (s string) {
	var err error
	if s, err = os.Getwd(); err != nil {
		return "noname"
	}
	return filepath.Base(s)
}
