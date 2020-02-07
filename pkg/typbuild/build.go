package typbuild

import (
	"context"
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuild/stdbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Build tool
type Build struct {
	commanders  []typcore.BuildCommander
	prebuilders []typcore.Prebuilder
	releaser    typcore.Releaser
}

// New return new instance of build
func New() *Build {
	return &Build{
		prebuilders: []typcore.Prebuilder{
			newStandardPrebuilder(),
		},
	}
}

// WithCommands to set command
func (b *Build) WithCommands(commanders ...typcore.BuildCommander) *Build {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithRelease to set releaser
func (b *Build) WithRelease(releaser typcore.Releaser) *Build {
	b.releaser = releaser
	return b
}

// WithPrebuild to set prebuilder
func (b *Build) WithPrebuild(prebuilders ...typcore.Prebuilder) *Build {
	b.prebuilders = append(b.prebuilders, prebuilders...)
	return b
}

// Releaser return the releaser
func (b *Build) Releaser() typcore.Releaser {
	return b.releaser
}

// BuildCommands to return command
func (b *Build) BuildCommands(bc *typcore.BuildContext) []*cli.Command {
	cmds := []*cli.Command{
		stdbuild.CmdBuild(bc),
		stdbuild.CmdClean(),
		stdbuild.CmdRun(bc),
		stdbuild.CmdTest(),
		stdbuild.CmdMock(bc),
		stdbuild.CmdRelease(bc),
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
