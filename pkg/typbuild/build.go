package typbuild

import (
	"context"
	"fmt"

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
	Prebuild(ctx context.Context, c *Context) error
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}

// New return new instance of build
func New() *Build {
	return &Build{
		prebuilders: []Prebuilder{&standardPrebuilder{}},
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
func (b *Build) BuildCommands(c *Context) []*cli.Command {
	cmds := []*cli.Command{
		b.cmdBuild(c),
		b.cmdClean(),
		b.cmdRun(c),
		b.cmdTest(),
		b.cmdMock(c),
		b.cmdRelease(c),
	}
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.BuildCommands(c)...)
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
func (b *Build) Prebuild(ctx context.Context, c *Context) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, c); err != nil {
			return
		}
	}
	return
}
