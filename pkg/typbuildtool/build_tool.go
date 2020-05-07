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
	BuildSequences []interface{}
	Utilities      []Utility

	SkipPrecondition bool

	IncludeBranch   bool
	IncludeCommitID bool
}

// Validate build
func (b *BuildTool) Validate() (err error) {
	if len(b.BuildSequences) < 1 {
		return errors.New("No build-sequence")
	}

	for _, module := range b.BuildSequences {
		if err = common.Validate(module); err != nil {
			return fmt.Errorf("BuildTool: %w", err)
		}
	}

	for _, utility := range b.Utilities {
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

	for _, task := range b.Utilities {
		cmds = append(cmds, task.Commands(c)...)
	}
	return cmds
}

// Precondition for this project
func (b *BuildTool) Precondition(c *PreconditionContext) (err error) {
	if b.SkipPrecondition {
		c.Info("Skip the preconditon")
		return
	}

	app := c.Core.App
	if configurer, ok := app.(typcfg.Configurer); ok {
		if err = typcfg.Write(DefaultConfigFile, configurer); err != nil {
			return
		}
	}

	if err = typcfg.Write(DefaultConfigFile, b); err != nil {
		return
	}

	if preconditioner, ok := app.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	typcfg.Load(DefaultConfigFile)

	return
}

// Configurations of Build-Tool
func (b *BuildTool) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range b.BuildSequences {
		if configurer, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	for _, utility := range b.Utilities {
		if configurer, ok := utility.(typcfg.Configurer); ok {
			cfgs = append(cfgs, configurer.Configurations()...)
		}
	}

	return
}

// RunBuildTool to run the build-tool
func (b *BuildTool) RunBuildTool(c *typcore.Context) (err error) {
	return createBuildToolCli(b, c).Run(os.Args)
}
