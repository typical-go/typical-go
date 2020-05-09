// Package typvar contains typical variable
package typvar

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
)
