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
	Name   = name()

	AppBin      = fmt.Sprintf("%s/%s", Layout.Bin, Name)
	AppMainPath = fmt.Sprintf("%s/%s", Layout.Cmd, Name)

	BuildTool         = "buildtool"
	BuildToolBin      = fmt.Sprintf("%s/%s-%s", Layout.Bin, Name, BuildTool)
	BuildToolMainPath = fmt.Sprintf("%s/%s-%s", Layout.Cmd, Name, BuildTool)

	Prebuilder         = "prebuilder"
	PrebuilderBin      = fmt.Sprintf("%s/%s-%s", Layout.Bin, Name, Prebuilder)
	PrebuilderMainPath = fmt.Sprintf("%s/%s-%s", Layout.Cmd, Name, Prebuilder)

	ContextFile  = "typical/context.go"
	ChecksumFile = Layout.Metadata + "/checksum"
)

func name() (s string) {
	var err error
	if s, err = os.Getwd(); err != nil {
		return "noname"
	}
	return filepath.Base(s)
}
