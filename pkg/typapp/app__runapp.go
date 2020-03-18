package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// RunApp to run the applciation
func (a *TypicalApp) RunApp(d *typcore.Descriptor) (err error) {
	return a.App(d).Run(os.Args)
}

// App is the cli app
func (a *TypicalApp) App(d *typcore.Descriptor) *cli.App {
	c := &Context{
		Descriptor: d,
		TypicalApp: a,
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	if entryPoint := a.EntryPoint(); entryPoint != nil {
		app.Action = c.ActionFunc(entryPoint)
	}
	app.Commands = a.Commands(c)
	return app
}
