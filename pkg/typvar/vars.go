// Package typvar contains typical variable
package typvar

import "time"

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

	// ConfigFile location
	ConfigFile = ".env"

	// TestTimeout duration
	TestTimeout = 25 * time.Second

	// TestCoverProfile location
	TestCoverProfile = "cover.out"

	// PrecondFile location
	PrecondFile = "typical/precond_DO_NOT_EDIT.go"
)
