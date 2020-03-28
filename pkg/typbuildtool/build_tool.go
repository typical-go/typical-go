package typbuildtool

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
)

// BuildTool is typical Build Tool for golang project
type BuildTool struct {
	modules   []interface{}
	utilities []Utility

	binFolder string
	cmdFolder string

	includeBranch   bool
	includeCommitID bool
}

// BuildSequences create new instance of BuildTool with build-sequence
func BuildSequences(modules ...interface{}) *BuildTool {
	return &BuildTool{
		modules:   modules,
		binFolder: DefaultBinFolder,
		cmdFolder: DefaultCmdFolder,
	}
}

// WithUtilities return build-tool with new utilities
func (b *BuildTool) WithUtilities(utilities ...Utility) *BuildTool {
	b.utilities = utilities
	return b
}

// WithBinFolder return BuildTool with new binFolder
func (b *BuildTool) WithBinFolder(binFolder string) *BuildTool {
	b.binFolder = binFolder
	return b
}

// WithCmdFolder return BuildTool with new cmdFolder
func (b *BuildTool) WithCmdFolder(cmdFolder string) *BuildTool {
	b.cmdFolder = cmdFolder
	return b
}

// CmdFolder of build-tool
func (b *BuildTool) CmdFolder() string {
	return b.cmdFolder
}

// BinFolder of build-tool
func (b *BuildTool) BinFolder() string {
	return b.binFolder
}

// Validate build
func (b *BuildTool) Validate() (err error) {
	if len(b.modules) < 1 {
		return errors.New("No build modules")
	}
	for _, module := range b.modules {
		if err = common.Validate(module); err != nil {
			return fmt.Errorf("BuildTool: %w", err)
		}
	}

	return
}

// Build task
func (b *BuildTool) Build(c *BuildContext) (dists []BuildDistribution, err error) {
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
func (b *BuildTool) Publish(pc *PublishContext) (err error) {
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
func (b *BuildTool) Release(rc *ReleaseContext) (files []string, err error) {
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
func (b *BuildTool) Clean(c *BuildContext) (err error) {
	for _, module := range b.modules {
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
func (b *BuildTool) Precondition(c *BuildContext) (err error) {
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
