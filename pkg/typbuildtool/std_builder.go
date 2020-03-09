package typbuildtool

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// StdBuilder is standard builder
type StdBuilder struct {
	prebuilders []Prebuilder
}

// NewBuilder return new instance of standard builder
func NewBuilder() *StdBuilder {
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
func (b *StdBuilder) Build(c *Context) (binary string, err error) {
	binary = fmt.Sprintf("%s/%s", c.BinFolder, c.Name) // TODO: move to context
	src := fmt.Sprintf("./%s/%s", c.CmdFolder, c.Name) // TODO: move to context
	ctx := c.Cli.Context

	if err = b.prebuild(c); err != nil {
		return
	}

	cmd := buildkit.NewGoBuild(binary, src).Command(ctx)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Build the project")
	return binary, cmd.Run()
}

func (b *StdBuilder) prebuild(c *Context) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(c); err != nil {
			return
		}
	}
	return
}
