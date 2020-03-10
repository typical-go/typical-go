package typbuildtool

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"
)

// StdBuilder is standard builder
type StdBuilder struct {
}

// NewBuilder return new instance of standard builder
func NewBuilder() *StdBuilder {
	return &StdBuilder{}
}

// Build the project
func (b *StdBuilder) Build(c *BuildContext) (dist BuildDistribution, err error) {
	binary := fmt.Sprintf("%s/%s", c.BinFolder, c.Name)
	srcDir := fmt.Sprintf("%s/%s", c.CmdFolder, c.Name)
	src := fmt.Sprintf("./%s/main.go", srcDir)
	ctx := c.Cli.Context

	if _, err = os.Stat(src); os.IsNotExist(err) {
		os.MkdirAll(srcDir, 0777)
		data := &tmpl.AppMainData{
			TypicalPackage: typcore.TypicalPackage(c.ProjectPackage),
		}
		if err = runnerkit.NewWriteTemplate(src, tmpl.AppMain, data).Run(ctx); err != nil {
			return
		}
	}

	cmd := buildkit.NewGoBuild(binary, src).Command(ctx)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return
	}

	return NewBuildDistribution(binary), nil
}
