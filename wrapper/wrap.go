package wrapper

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"go/build"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/typvar"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

const (
	projectPkgVar = "github.com/typical-go/typical-go/pkg/typvar.ProjectPkg"
	typicalTmpVar = "github.com/typical-go/typical-go/pkg/typvar.TypicalTmp"
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
func Wrap(c *Context) (err error) {

	if c.ProjectPkg == "" {
		if c.ProjectPkg, err = retrieveProjectPackage(); err != nil {
			return
		}
	}

	typvar.TypicalTmp = c.TypicalTmp
	typvar.ProjectPkg = c.ProjectPkg

	// NOTE: create tmp folder if not exist
	os.MkdirAll(c.TypicalTmp+"/build-tool", os.ModePerm)
	os.MkdirAll(c.TypicalTmp+"/bin", os.ModePerm)

	gitignore := ".gitignore"
	if _, err = os.Stat(gitignore); os.IsNotExist(err) {
		c.Infof("Generate %s", gitignore)
		if err = typtmpl.WriteFile(gitignore, 0777, &typtmpl.GitIgnore{}); err != nil {
			return
		}
	}

	typicalw := "typicalw"
	if _, err = os.Stat(typicalw); os.IsNotExist(err) {
		c.Infof("Generate %s", typicalw)
		if err = typtmpl.WriteFile(typicalw, 0777, &typtmpl.Typicalw{
			TypicalSource: "github.com/typical-go/typical-go/cmd/typical-go",
			TypicalTmp:    c.TypicalTmp,
			ProjectPkg:    c.ProjectPkg,
		}); err != nil {
			return
		}
	}

	// checksumPath := fmt.Sprintf("%s/checksum", c.TypicalTmp)
	// buildTool := c.TypicalTmp + "/bin/build-tool"
	// srcPath := c.TypicalTmp + "/build-tool/main.go"

	build := typvar.GetBuild()

	descriptorPkg := fmt.Sprintf("%s/%s", c.ProjectPkg, c.DescriptorPkg)

	var checksum *Checksum
	if checksum, err = CreateChecksum(c.DescriptorPkg); err != nil {
		return
	}

	if _, err = os.Stat(build.Binary); os.IsNotExist(err) || !checksum.IsSame(build.Checksum) {
		c.Info("Update checksum")
		if err = checksum.Save(build.Checksum); err != nil {
			return
		}

		if _, err = os.Stat(build.Source); os.IsNotExist(err) {
			c.Infof("Generate build-tool main source: %s", build.Source)
			if err = typtmpl.WriteFile(build.Source, 0777, &typtmpl.BuildToolMain{
				DescPkg: descriptorPkg,
			}); err != nil {
				return
			}
		}

		c.Info("Build the build-tool")

		gobuild := &buildkit.GoBuild{
			Out:    build.Binary,
			Source: build.Source,
			Ldflags: []string{
				buildkit.BuildVar(projectPkgVar, c.ProjectPkg),
				buildkit.BuildVar(typicalTmpVar, c.TypicalTmp),
			},
		}

		cmd := gobuild.Command()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		return cmd.Run(c.Ctx)
	}
	return
}

// TODO: using go list -m to get package name
func retrieveProjectPackage() (pkg string, err error) {
	var (
		root string
		f    *os.File
	)

	if root, err = os.Getwd(); err != nil {
		return
	}

	// go.mod is not exist. Check if the project sit in $GOPATH
	if f, err = os.Open(root + "/go.mod"); err != nil {
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			pkg = root[len(gopath):]
		} else {
			err = errors.New("RetrieveProjectPackage: go.mod is missing and the project not in $GOPATH")
		}
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			pkg = strings.TrimSpace(line[6:])
			return
		}
	}

	err = errors.New("RetrieveProjectPackage: go.mod doesn't contain module")
	return
}
