package typgo

import "github.com/fatih/color"

// only available in project scope as supplied when compile using typgo.GoBuild or ldflags
var (
	// ProjectName of application.
	ProjectName string
	// ProjectVersion of applicatoin.
	ProjectVersion string
)

// only available in build-tool scope
var (
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string
	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string
)

// template
const (
	appHelpTemplate = `{{range .VisibleCommands}}{{if not .HideHelp}}{{ "\t"}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}
{{end}}{{end}}`
	subcommandHelpTemplate = `{{.Usage}}

Subtasks:{{range .VisibleCategories}}{{if .Name}}{{.Name}}:{{range .VisibleCommands}}
		{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
	{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
`
)

// ColorSet color set
var ColorSet = struct {
	Project color.Attribute
	Task    color.Attribute
	Bash    color.Attribute
	Warn    color.Attribute
}{
	Project: color.FgHiCyan,
	Task:    color.FgCyan,
	Bash:    color.FgGreen,
	Warn:    color.FgYellow,
}
