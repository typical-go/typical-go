package typgo

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

var (
	_ AppLauncher       = (*Descriptor)(nil)
	_ BuildToolLauncher = (*Descriptor)(nil)

	_ Utility        = (*Descriptor)(nil)
	_ Preconditioner = (*Descriptor)(nil)
)

type (

	// Descriptor describe the project
	Descriptor struct {

		// Name of the project (OPTIONAL).
		// It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		Name string

		// Description of the project (OPTIONAL).
		Description string

		// Version of the project (OPTIONAL).
		// By default it is 0.0.1
		Version string

		// App of the project (MANDATORY).
		App Runner

		BuildSequences []interface{}
		Utility        Utility
		Layouts        []string

		SkipPrecond bool
	}

	// Runner responsible to run the application
	Runner interface {
		Run(*Descriptor) error
	}
)

// LaunchApp to launch the app
func (d *Descriptor) LaunchApp() (err error) {
	if err = d.Validate(); err != nil {
		return
	}
	return d.App.Run(d)
}

// LaunchBuildTool to launch the build tool
func (d *Descriptor) LaunchBuildTool() (err error) {
	if err := d.Validate(); err != nil {
		return err
	}

	appDirs, appFiles := WalkLayout(d.Layouts)

	cli := createBuildToolCli(&Context{
		Descriptor: d,
		AppDirs:    appDirs,
		AppFiles:   appFiles,
	})
	return cli.Run(os.Args)
}

// Validate context
func (d *Descriptor) Validate() (err error) {
	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if !ValidateName(d.Name) {
		return errors.New("Descriptor: bad name")
	}

	if err = common.Validate(d.App); err != nil {
		return fmt.Errorf("Descriptor: App: %w", err)
	}

	if len(d.BuildSequences) < 1 {
		return errors.New("Descriptor:No build-sequence")
	}

	for _, module := range d.BuildSequences {
		if err = common.Validate(module); err != nil {
			return err
		}
	}

	return
}

// ValidateName to validate valid descriptor name
func ValidateName(name string) bool {
	if name == "" {
		return false
	}

	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return false
	}
	return true
}

// Commands to return command
func (d *Descriptor) Commands(c *Context) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(c),
		cmdRun(c),
		cmdPublish(c),
		cmdClean(c),
	}

	if d.Utility != nil {
		for _, cmd := range d.Utility.Commands(c) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

// Precondition for this project
func (d *Descriptor) Precondition(c *PrecondContext) (err error) {
	if d.SkipPrecond {
		c.Info("Skip the preconditon")
		return
	}

	app := c.App
	if configurer, ok := app.(typcfg.Configurer); ok {
		if err = typcfg.Write(typvar.ConfigFile, configurer); err != nil {
			return
		}
	}

	if preconditioner, ok := app.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	typcfg.Load(typvar.ConfigFile)

	return
}
