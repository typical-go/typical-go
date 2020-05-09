package typbuild

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.Runner    = (*BuildTool)(nil)
	_ typcfg.Configurer = (*BuildTool)(nil)
	_ Utility           = (*BuildTool)(nil)
	_ Preconditioner    = (*BuildTool)(nil)
)

// BuildTool is typical Build Tool for golang project
type BuildTool struct {
	BuildSequences []interface{}
	Utilities      []Utility
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

	for _, utility := range b.Utilities {
		if err = common.Validate(utility); err != nil {
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

	for _, task := range b.Utilities {
		cmds = append(cmds, task.Commands(c)...)
	}
	return cmds
}

// Precondition for this project
func (b *BuildTool) Precondition(c *PrecondContext) (err error) {
	if b.SkipPrecond {
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

// Run the build-tool
func (b *BuildTool) Run(d *typcore.Descriptor) (err error) {

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

// WalkLayout return dirs and files
func WalkLayout(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if isGoSource(path) {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

func isGoSource(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}
