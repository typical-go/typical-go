package typbuildtool

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.BuildTool = (*BuildTool)(nil)
	_ typcfg.Configurer = (*BuildTool)(nil)
	_ Utility           = (*BuildTool)(nil)
	_ Preconditioner    = (*BuildTool)(nil)
)

// BuildTool is typical Build Tool for golang project
type BuildTool struct {
	buildSequences []interface{}
	utilities      []Utility

	binFolder string // TODO: move to context
	cmdFolder string // TODO: move to context

	configFile         string
	enablePrecondition bool

	includeBranch   bool
	includeCommitID bool
}

// BuildSequences create new instance of BuildTool with build-sequence
func BuildSequences(buildSequences ...interface{}) *BuildTool {
	return &BuildTool{
		buildSequences:     buildSequences,
		binFolder:          DefaultBinFolder,
		cmdFolder:          DefaultCmdFolder,
		configFile:         DefaultConfigFile,
		enablePrecondition: DefaultEnablePrecondition,
	}
}

// Utilities return build-tool with new utilities
func (b *BuildTool) Utilities(utilities ...Utility) *BuildTool {
	b.utilities = utilities
	return b
}

// EnablePrecondition define whether execute precondition or not. By default if true
func (b *BuildTool) EnablePrecondition(enablePrecondition bool) *BuildTool {
	b.enablePrecondition = enablePrecondition
	return b
}

// Validate build
func (b *BuildTool) Validate() (err error) {
	if len(b.buildSequences) < 1 {
		return errors.New("No build-sequence")
	}

	for _, module := range b.buildSequences {
		if err = common.Validate(module); err != nil {
			return fmt.Errorf("BuildTool: %w", err)
		}
	}

	for _, utility := range b.utilities {
		if err = common.Validate(utility); err != nil {
			return fmt.Errorf("BuildTool: %w", err)
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

	for _, task := range b.utilities {
		cmds = append(cmds, task.Commands(c)...)
	}
	return cmds
}

// Precondition for this project
func (b *BuildTool) Precondition(c *PreconditionContext) (err error) {
	if !b.enablePrecondition {
		c.Info("Skip the preconditon")
		return
	}

	app := c.Core.App
	if configurer, ok := app.(typcfg.Configurer); ok {
		if err = typcfg.Write(b.configFile, configurer); err != nil {
			return
		}
	}

	if err = typcfg.Write(b.configFile, b); err != nil {
		return
	}

	if preconditioner, ok := app.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	typcfg.Load(b.configFile)

	return
}

// Configurations of Build-Tool
func (b *BuildTool) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range b.buildSequences {
		if configurer, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	for _, utility := range b.utilities {
		if configurer, ok := utility.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	return
}

// RunBuildTool to run the build-tool
func (b *BuildTool) RunBuildTool(c *typcore.Context) (err error) {
	return b.cli(c).Run(os.Args)
}

func (b *BuildTool) cli(core *typcore.Context) *cli.App {
	app := cli.NewApp()
	app.Name = core.Name
	app.Usage = "Build-Tool"
	app.Description = core.Description
	app.Version = core.Version

	app.Before = func(cli *cli.Context) (err error) {
		return b.Precondition(&PreconditionContext{
			Core: core,
			Ctx:  cli.Context,
		})
	}
	app.Commands = b.Commands(&Context{
		Core:      core,
		BuildTool: b,
	})

	return app
}
