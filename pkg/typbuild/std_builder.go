package typbuild

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// StdBuilder is standard builder
type StdBuilder struct {
	prebuilders []Prebuilder
}

// New return new instance of standard builder
func New() *StdBuilder {
	return &StdBuilder{
		prebuilders: []Prebuilder{
			&stdPrebuilder{},
		},
	}
}

// AppendPrebuilder return StdBuilder with appended prebuilder
func (b *StdBuilder) AppendPrebuilder(prebuilders ...Prebuilder) *StdBuilder {
	b.prebuilders = append(b.prebuilders, prebuilders...)
	return b
}

// Build the project
func (b *StdBuilder) Build(ctx context.Context, c *Context) (binary string, err error) {
	binary = fmt.Sprintf("%s/%s", c.BinFolder, c.Name) // TODO: move to context
	src := fmt.Sprintf("./%s/%s", c.CmdFolder, c.Name) // TODO: move to context
	if err = b.prebuild(ctx, c); err != nil {
		return
	}

	cmd := buildkit.NewGoBuild(binary, src).Command(ctx)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Build the project")
	return binary, cmd.Run()
}

func (b *StdBuilder) prebuild(ctx context.Context, c *Context) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, c); err != nil {
			return
		}
	}
	return
}
