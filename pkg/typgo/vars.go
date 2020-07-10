package typgo

import (
	"fmt"
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

	TmpBin string
	TmpSrc string

	ExclMsgPrefix = []string{
		"merge", "bump", "revision", "generate", "wip",
	}
)

var (
	appHelpTemplate = `Typical Build

Usage:

{{"\t"}}./typicalw <command> [argument]

The commands are:
{{range .Commands}}
{{if not .HideHelp}}{{ "\t"}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{end}}{{end}}

Use "./typicalw help <topic>" for more information about that topic
`
	subcommandHelpTemplate = `{{.Usage}}

Usage:

	{{.Name}} [command]
	
Commands:{{range .VisibleCategories}}
{{if .Name}}{{.Name}}:{{range .VisibleCommands}}
		{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
	{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
	
{{if .VisibleFlags}} 
Options:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`
)

// AppBin path
func AppBin(name string) string {
	return fmt.Sprintf("%s/%s", BinFolder, name)
}

// Init vars
func Init() error {
	TmpBin = fmt.Sprintf("%s/bin", TypicalTmp)
	TmpSrc = fmt.Sprintf("%s/src", TypicalTmp)
	return nil
}
