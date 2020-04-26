package typbuildtool

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.BuildTool = (*BuildTool)(nil)
	_ typcfg.Configurer = (*BuildTool)(nil)

	_ Utility        = (*BuildTool)(nil)
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

// Commands to return command
func (b *BuildTool) Commands(c *Context) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(c),
		cmdRun(c),
		b.cmdPublish(c),
		b.cmdClean(c),
	}

	for _, task := range b.utilities {
		cmds = append(cmds, task.Commands(c)...)
	}
	return cmds
}
