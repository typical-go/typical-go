package typbuildtool

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// TypicalBuildTool is typical Build Tool for golang project
type TypicalBuildTool struct {
	commanders []Commander
	builder    Builder
	tester     Tester
	mocker     Mocker
	releaser   Releaser
	publishers []Publisher

	binFolder string
	cmdFolder string

	includeBranch   bool
	includeCommitID bool
}

// New return new instance of build
func New() *TypicalBuildTool {
	return &TypicalBuildTool{
		builder:   NewBuilder(),
		tester:    NewTester(),
		mocker:    NewMocker(),
		releaser:  NewReleaser(),
		binFolder: "bin",
		cmdFolder: "cmd",
	}
}

// AppendCommanders to return BuildTool with appended commander
func (b *TypicalBuildTool) AppendCommanders(commanders ...Commander) *TypicalBuildTool {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithBuilder return  BuildTool with new builder
func (b *TypicalBuildTool) WithBuilder(builder Builder) *TypicalBuildTool {
	b.builder = builder
	return b
}

// WithReleaser return BuildTool with new releaser
func (b *TypicalBuildTool) WithReleaser(releaser Releaser) *TypicalBuildTool {
	b.releaser = releaser
	return b
}

// WithPublishers return BuildTool with new publishers
func (b *TypicalBuildTool) WithPublishers(publishers ...Publisher) *TypicalBuildTool {
	b.publishers = publishers
	return b
}

// WithMocker return BuildTool with new mocker
func (b *TypicalBuildTool) WithMocker(mocker Mocker) *TypicalBuildTool {
	b.mocker = mocker
	return b
}

// WithTester return BuildTool with new tester
func (b *TypicalBuildTool) WithTester(tester Tester) *TypicalBuildTool {
	b.tester = tester
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

// Validate build
func (b *TypicalBuildTool) Validate() (err error) {

	if err = common.Validate(b.builder); err != nil {
		return fmt.Errorf("BuildTool: Builder: %w", err)
	}

	if err = common.Validate(b.mocker); err != nil {
		return fmt.Errorf("BuildTool: Mocker: %w", err)
	}

	if err = common.Validate(b.tester); err != nil {
		return fmt.Errorf("BuildTool: Tester: %w", err)
	}

	if err = common.Validate(b.releaser); err != nil {
		return fmt.Errorf("BuildTool: Releaser: %w", err)
	}

	return
}

// Build task
func (b *TypicalBuildTool) Build(c *Context) (dist BuildDistribution, err error) {
	return b.builder.Build(c)
}

// Publish the project
func (b *TypicalBuildTool) Publish(pc *PublishContext) (err error) {
	for _, publisher := range b.publishers {
		if err = publisher.Publish(pc); err != nil {
			return
		}
	}
	return
}

// Release the project
func (b *TypicalBuildTool) Release(rc *ReleaseContext) ([]string, error) {
	return b.releaser.Release(rc)
}

// Clean the project
func (b *TypicalBuildTool) Clean(c *Context) (err error) {

	if cleaner, ok := b.builder.(Cleaner); ok {
		cleaner.Clean(c)
	}

	if c.Cli.Bool("short") {
		os.Remove(typcore.BuildToolBin(c.TempFolder))
	} else {
		c.Infof("Remove All: %s", c.TempFolder)
		if err := os.RemoveAll(c.TempFolder); err != nil {
			c.Error(err.Error())
		}
	}
	return
}

// Test the project
func (b *TypicalBuildTool) Test(c *Context) error {
	if b.tester == nil {
		panic("TypicalBuildTool: missing tester")
	}
	return b.tester.Test(c)
}
