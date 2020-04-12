package typbuildtool

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	_ typcore.BuildTool = (*BuildTool)(nil)
	_ typcfg.Configurer = (*BuildTool)(nil)

	_ Utility        = (*BuildTool)(nil)
	_ Runner         = (*BuildTool)(nil)
	_ Tester         = (*BuildTool)(nil)
	_ Cleaner        = (*BuildTool)(nil)
	_ Releaser       = (*BuildTool)(nil)
	_ Publisher      = (*BuildTool)(nil)
	_ Preconditioner = (*BuildTool)(nil)
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

// BinFolder return BuildTool with new binFolder
func (b *BuildTool) BinFolder(binFolder string) *BuildTool {
	b.binFolder = binFolder
	return b
}

// CmdFolder return BuildTool with new cmdFolder
func (b *BuildTool) CmdFolder(cmdFolder string) *BuildTool {
	b.cmdFolder = cmdFolder
	return b
}

// ConfigFile define path to store config
func (b *BuildTool) ConfigFile(configFile string) *BuildTool {
	b.configFile = configFile
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

// Run task
func (b *BuildTool) Run(c *BuildContext) (err error) {
	for _, module := range b.buildSequences {
		if runner, ok := module.(Runner); ok {
			if err = runner.Run(c); err != nil {
				return
			}
		}
	}
	return
}

// Publish the project
func (b *BuildTool) Publish(pc *PublishContext) (err error) {
	for _, module := range b.buildSequences {
		if publisher, ok := module.(Publisher); ok {
			if err = publisher.Publish(pc); err != nil {
				return
			}
		}
	}
	return
}

// Release the project
func (b *BuildTool) Release(rc *ReleaseContext) (files []string, err error) {
	for _, module := range b.buildSequences {
		if releaser, ok := module.(Releaser); ok {
			var files1 []string
			if files1, err = releaser.Release(rc); err != nil {
				return
			}
			files = append(files, files1...)
		}
	}
	return
}

// Clean the project
func (b *BuildTool) Clean(c *BuildContext) (err error) {
	for _, module := range b.buildSequences {
		if cleaner, ok := module.(Cleaner); ok {
			if err = cleaner.Clean(c); err != nil {
				return
			}
		}
	}

	c.Infof("Remove All: %s", c.TypicalTmp)
	if err := os.RemoveAll(c.TypicalTmp); err != nil {
		c.Warn(err.Error())
	}

	return
}

// Test the project
func (b *BuildTool) Test(c *BuildContext) (err error) {
	for _, module := range b.buildSequences {
		if tester, ok := module.(Tester); ok {
			if err = tester.Test(c); err != nil {
				return
			}
		}
	}
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
