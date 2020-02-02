package typcore

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

// NewBuild return new instance of build
func NewBuild() *Build {
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

// Releaser return the releaser
func (b *Build) Releaser() Releaser {
	return b.releaser
}

// BuildCommands to return command
func (b *Build) BuildCommands(bc *BuildContext) (cmds []*cli.Command) {
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.BuildCommands(bc)...)
	}
	return
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
func (b *Build) Prebuild(ctx context.Context, bc *BuildContext) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, bc); err != nil {
			return
		}
	}
	return
}
