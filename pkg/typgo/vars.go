package typgo

import "github.com/fatih/color"

var (
	// ProjectName of application. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=PROJECT-NAME`
	ProjectName string
	// ProjectVersion of applicatoin. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=PROJECT-NAME`
	ProjectVersion string
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string
	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string

	appHelpTemplate = `Tasks:
{{range .VisibleCommands}}{{if not .HideHelp}}{{ "\t"}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}
{{end}}{{end}}`
	subcommandHelpTemplate = `{{.Usage}}

Subtasks:{{range .VisibleCategories}}{{if .Name}}{{.Name}}:{{range .VisibleCommands}}
		{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
	{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
`
)

const (
	ProjectNameColor = color.FgCyan
	PrepareColor     = color.FgHiBlue
	TaskNameColor    = color.FgBlue
	BashColor        = color.FgGreen
)
