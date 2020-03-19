package typbuildtool

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
)

// TypicalBuildTool is typical Build Tool for golang project
type TypicalBuildTool struct {
	modules    []interface{}
	commanders []Commander

	binFolder string
	cmdFolder string
	tmpFolder string

	includeBranch   bool
	includeCommitID bool
}

// New return new instance of build
func New(modules ...interface{}) *TypicalBuildTool {
	return &TypicalBuildTool{
		modules:   modules,
		binFolder: DefaultBinFolder,
		cmdFolder: DefaultCmdFolder,
		tmpFolder: DefaultTmpFolder,
	}
}

// WithCommanders return build-tool with commanders
func (b *TypicalBuildTool) WithCommanders(commanders ...Commander) *TypicalBuildTool {
	b.commanders = commanders
	return b
}

// WithBinFolder return BuildTool with new binFolder
func (b *TypicalBuildTool) WithBinFolder(binFolder string) *TypicalBuildTool {
	b.binFolder = binFolder
	return b
}

// WithCmdFolder return BuildTool with new cmdFolder
func (b *TypicalBuildTool) WithCmdFolder(cmdFolder string) *TypicalBuildTool {
	b.cmdFolder = cmdFolder
	return b
}

// CmdFolder of build-tool
func (b *TypicalBuildTool) CmdFolder() string {
	return b.cmdFolder
}

// BinFolder of build-tool
func (b *TypicalBuildTool) BinFolder() string {
	return b.binFolder
}

// TmpFolder of build-tool
func (b *TypicalBuildTool) TmpFolder() string {
	return b.tmpFolder
}

// Validate build
func (b *TypicalBuildTool) Validate() (err error) {
	for _, module := range b.modules {
		if err = common.Validate(module); err != nil {
			return fmt.Errorf("BuildTool: %w", err)
		}
	}

	return
}

// Build task
func (b *TypicalBuildTool) Build(c *BuildContext) (dists []BuildDistribution, err error) {
	c.Info("Build the project")
	for _, module := range b.modules {
		if builder, ok := module.(Builder); ok {
			var dists1 []BuildDistribution
			if dists1, err = builder.Build(c); err != nil {
				return
			}
			dists = append(dists, dists1...)
		}
	}
	return
}

// Publish the project
func (b *TypicalBuildTool) Publish(pc *PublishContext) (err error) {
	for _, module := range b.modules {
		if publisher, ok := module.(Publisher); ok {
			if err = publisher.Publish(pc); err != nil {
				return
			}
		}
	}
	return
}

// Release the project
func (b *TypicalBuildTool) Release(rc *ReleaseContext) (files []string, err error) {
	rc.Info("Release the project")
	for _, module := range b.modules {
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
func (b *TypicalBuildTool) Clean(c *BuildContext) (err error) {
	for _, module := range b.modules {
		if cleaner, ok := module.(Cleaner); ok {
			if err = cleaner.Clean(c); err != nil {
				return
			}
		}
	}

	// TODO: move to module
	if c.Cli.Bool("short") {
		os.Remove(b.tmpFolder + "/build-tool")
	} else {
		c.Infof("Remove All: %s", b.tmpFolder)
		if err := os.RemoveAll(b.tmpFolder); err != nil {
			c.Error(err.Error())
		}
	}
	return
}

// Test the project
func (b *TypicalBuildTool) Test(c *BuildContext) (err error) {
	c.Info("Test the project")
	for _, module := range b.modules {
		if tester, ok := module.(Tester); ok {
			if err = tester.Test(c); err != nil {
				return
			}
		}
	}
	return
}

// Precondition for this project
func (b *TypicalBuildTool) Precondition(c *Context) (err error) {
	if preconditioner, ok := c.App.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-App: %w", err)
		}
	}

	if preconditioner, ok := c.ConfigManager.(Preconditioner); ok {
		if err = preconditioner.Precondition(c); err != nil {
			return fmt.Errorf("Precondition-Config-Manager: %w", err)
		}
	}

	return
}
