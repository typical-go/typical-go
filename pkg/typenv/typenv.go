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

	Readme      = "README.md"
	ProjectName = projectName()

	AppBin      = fmt.Sprintf("%s/%s", Layout.Bin, ProjectName)
	AppMainPath = fmt.Sprintf("%s/%s", Layout.Cmd, ProjectName)

	BuildTool         = "buildtool"
	BuildToolBin      = fmt.Sprintf("%s/%s-%s", Layout.Bin, ProjectName, BuildTool)
	BuildToolMainPath = fmt.Sprintf("%s/%s-%s", Layout.Cmd, ProjectName, BuildTool)

	ContextFile  = "typical/context.go"
	ChecksumFile = Layout.Metadata + "/checksum"
)

func projectName() (s string) {
	var err error
	if s, err = os.Getwd(); err != nil {
		return "noname"
	}
	return filepath.Base(s)
}
