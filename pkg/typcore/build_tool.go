package typcore

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

var (
	_ Runner         = (*BuildTool)(nil)
	_ Utility        = (*BuildTool)(nil)
	_ Preconditioner = (*BuildTool)(nil)
)

// BuildTool is typical Build Tool for golang project
type BuildTool struct {
	BuildSequences []interface{}
	Utility        Utility
	Layouts        []string

	SkipPrecond bool
}

// Validate build
func (b *BuildTool) Validate() (err error) {
	if len(b.BuildSequences) < 1 {
		return errors.New("No build-sequence")
	}

	for _, module := range b.BuildSequences {
		if err = common.Validate(module); err != nil {
			return err
		}
	}

	return
}

// Commands to return command
func (b *BuildTool) Commands(c *Context) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(c),
		cmdRun(c),
		cmdPublish(c),
		cmdClean(c),
	}

	if b.Utility != nil {
		for _, cmd := range b.Utility.Commands(c) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

// Precondition for this project
func (b *BuildTool) Precondition(c *PrecondContext) (err error) {
	if b.SkipPrecond {
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

// Run the build-tool
func (b *BuildTool) Run(d *Descriptor) (err error) {

	if err := d.Validate(); err != nil {
		return err
	}

	appDirs, appFiles := WalkLayout(b.Layouts)

	cli := createBuildToolCli(b, &Context{
		Descriptor: d,
		AppDirs:    appDirs,
		AppFiles:   appFiles,
		BuildTool:  b,
	})
	return cli.Run(os.Args)
}
