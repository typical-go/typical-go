package typbuild

import (
	"context"
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Build tool
type Build struct {
	commanders  []BuildCommander
	prebuilders []Prebuilder
	releaser    Releaser
}

// Releaser responsible to release
type Releaser interface {
	BuildRelease(ctx context.Context, name, tag string, changeLogs []string, alpha bool) (binaries []string, err error)
	Publish(ctx context.Context, name, tag string, changeLogs, binaries []string, alpha bool) (err error)
	Tag(ctx context.Context, version string, alpha bool) (tag string, err error)
	Validate() error
}

// Prebuilder responsible to prebuild task
type Prebuilder interface {
	Prebuild(ctx context.Context, bc *typcore.BuildContext) error
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *typcore.BuildContext) []*cli.Command
}

// New return new instance of build
func New() *Build {
	return &Build{
		prebuilders: []Prebuilder{
			newStandardPrebuilder(),
		},
	}
}

// WithCommands to set command
func (b *Build) WithCommands(commanders ...BuildCommander) *Build {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithRelease to set releaser
func (b *Build) WithRelease(releaser Releaser) *Build {
	b.releaser = releaser
	return b
}

// WithPrebuild to set prebuilder
func (b *Build) WithPrebuild(prebuilders ...Prebuilder) *Build {
	b.prebuilders = append(b.prebuilders, prebuilders...)
	return b
}

// BuildCommands to return command
func (b *Build) BuildCommands(bc *typcore.BuildContext) []*cli.Command {
	cmds := []*cli.Command{
		b.cmdBuild(bc),
		b.cmdClean(),
		b.cmdRun(bc),
		b.cmdTest(),
		b.cmdMock(bc),
		b.cmdRelease(bc),
	}
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.BuildCommands(bc)...)
	}
	return cmds
}

// Validate build
func (b *Build) Validate() (err error) {
	if b.releaser != nil {
		if err = b.releaser.Validate(); err != nil {
			return fmt.Errorf("Build: Releaser: %w", err)
		}
	}
	return
}

// Prebuild process
func (b *Build) Prebuild(ctx context.Context, bc *typcore.BuildContext) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, bc); err != nil {
			return
		}
	}
	return
}
