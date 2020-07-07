package typgo

import (
	"fmt"
	"os"
)

var (
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string

	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string

	// ReleaseFolder location
	ReleaseFolder = "release"

	// BinFolder location
	BinFolder = "bin"

	// CmdFolder location
	CmdFolder = "cmd"

	TmpBin        string
	TmpSrc        string
	BuildChecksum string
	BuildToolSrc  string
	BuildToolBin  string

	ExclMsgPrefix = []string{
		"merge", "bump", "revision", "generate", "wip",
	}
)

// Precond path
func Precond(name string) string {
	return fmt.Sprintf("%s/%s/precond_DO_NOT_EDIT.go", CmdFolder, name)
}

// AppBin path
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", BinFolder, name)
}

// Init vars
func Init() error {
	TmpBin = fmt.Sprintf("%s/bin", TypicalTmp)
	TmpSrc = fmt.Sprintf("%s/src", TypicalTmp)
	BuildChecksum = fmt.Sprintf("%s/checksum", TypicalTmp)
	BuildToolSrc = fmt.Sprintf("%s/build-tool", TmpSrc)
	BuildToolBin = fmt.Sprintf("%s/build-tool", TmpBin)
	return nil
}

// Wrap vars
func Wrap(typicalTmp, projectPkg string) error {
	TypicalTmp = typicalTmp
	ProjectPkg = projectPkg
	Init()

	os.MkdirAll(BuildToolSrc, 0777)
	os.MkdirAll(TmpBin, 0777)
	return nil
}
