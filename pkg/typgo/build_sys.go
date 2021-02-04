package typgo

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// Descriptor describe the project
	Descriptor struct {
		ProjectName    string // By default is same with project folder. Only allowed characters(a-z,A-Z), underscore or dash.
		ProjectVersion string // By default it is 0.0.1
		EnvLoader      EnvLoader
		Tasks          []Tasker
	}
)

// Start typical build-tool
func Start(d *Descriptor) {
	if err := Cli(d).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Action to return cli func from action
func (d *Descriptor) Action(action Action) func(*cli.Context) error {
	return func(cliCtx *cli.Context) error {
		if action == nil {
			return nil
		}
		return action.Execute(&Context{
			Context:    cliCtx,
			Descriptor: d,
		})
	}
}

// Execute action
func (d *Descriptor) Execute(action Action, cliCtx *cli.Context) error {
	return action.Execute(&Context{
		Context:    cliCtx,
		Descriptor: d,
	})
}
