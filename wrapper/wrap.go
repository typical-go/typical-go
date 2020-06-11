package wrapper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

const (
	projectPkgVar = "github.com/typical-go/typical-go/pkg/typvar.ProjectPkg"
	typicalTmpVar = "github.com/typical-go/typical-go/pkg/typvar.TypicalTmp"
	gitignore     = ".gitignore"
	typicalw      = "typicalw"
)

// Context of wrapper
type Context struct {
	*typgo.Descriptor
	typlog.Logger

	Ctx context.Context

	TypicalTmp    string
	ProjectPkg    string
	DescriptorPkg string
}

// Wrap the project
func Wrap(c *Context) error {
	if c.ProjectPkg == "" {
		var stderr strings.Builder
		var stdout strings.Builder

		cmd := execkit.Command{
			Name:   "go",
			Args:   []string{"list", "-m"},
			Stdout: &stdout,
			Stderr: &stderr,
		}

		if err := cmd.Run(c.Ctx); err != nil {
			return errors.New(stderr.String())
		}

		c.ProjectPkg = strings.TrimSpace(stdout.String())
	}

	typvar.Wrap(c.TypicalTmp, c.ProjectPkg)

	if _, err := os.Stat(gitignore); os.IsNotExist(err) {
		c.Infof("Generate %s", gitignore)
		if err = typtmpl.WriteFile(gitignore, 0777, &typtmpl.GitIgnore{}); err != nil {
			return err
		}
	}

	if _, err := os.Stat(typicalw); os.IsNotExist(err) {
		c.Infof("Generate %s", typicalw)
		if err = typtmpl.WriteFile(typicalw, 0777, &typtmpl.Typicalw{
			TypicalSource: "github.com/typical-go/typical-go/cmd/typical-go",
			TypicalTmp:    c.TypicalTmp,
			ProjectPkg:    c.ProjectPkg,
		}); err != nil {
			return err
		}
	}

	descriptorPkg := fmt.Sprintf("%s/%s", c.ProjectPkg, c.DescriptorPkg)

	checksum, err := CreateChecksum(c.DescriptorPkg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(typvar.BuildToolSrc + "/main.go"); os.IsNotExist(err) {
		if err = typtmpl.WriteFile(typvar.BuildToolSrc+"/main.go", 0777, &typtmpl.BuildToolMain{
			DescPkg: descriptorPkg,
		}); err != nil {
			return err
		}
	}

	if _, err = os.Stat(typvar.BuildToolBin); os.IsNotExist(err) || !checksum.IsSame(typvar.BuildChecksum) {
		if err = checksum.Save(typvar.BuildChecksum); err != nil {
			return err
		}

		c.Info("Build the build-tool")
		gobuild := &buildkit.GoBuild{
			Out:    typvar.BuildToolBin,
			Source: "./" + typvar.BuildToolSrc,
			Ldflags: []string{
				buildkit.BuildVar(projectPkgVar, c.ProjectPkg),
				buildkit.BuildVar(typicalTmpVar, c.TypicalTmp),
			},
		}

		return gobuild.Run(c.Ctx)
	}
	return nil
}
