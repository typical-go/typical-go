package typgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate = `Typical Build

Usage:

{{"\t"}}./typicalw <command> [argument]

The commands are:
{{range .Commands}}
{{if not .HideHelp}}{{ "\t"}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{end}}{{end}}

Use "./typicalw help <topic>" for more information about that topic
`

	cli.SubcommandHelpTemplate = `{{.Usage}}

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

}

func launchBuild(d *Descriptor) error {
	if err := typvar.Init(); err != nil {
		return fmt.Errorf("init-var: %w", err)
	}

	buildCli := createBuildCli(d)
	if err := buildCli.Prebuild(); err != nil {
		return err
	}

	cmds, err := buildCli.commands()
	if err != nil {
		return err
	}

	app := cli.NewApp()
	app.Name = "./typicalw"
	app.Usage = "./tyicalw"
	app.After = buildCli.ActionFn("AFTER_BUILD", afterBuild)
	app.Commands = cmds

	return app.Run(os.Args)
}

func afterBuild(c *Context) (err error) {
	store := c.BuildCli.ASTStore
	b, _ := json.MarshalIndent(store.Annots, "", "\t")
	if err = ioutil.WriteFile(fmt.Sprintf("%s/annots.json", typvar.TypicalTmp), b, 0777); err != nil {
		return
	}
	return
}
