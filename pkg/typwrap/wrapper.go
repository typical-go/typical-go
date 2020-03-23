package typwrap

import (
	"go/build"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// TypicalWrapper responsible to wrap the typical project
type TypicalWrapper struct{}

// New instance of TypicalWrapper
func New() *TypicalWrapper {
	return &TypicalWrapper{}
}

// Wrap the project
func (*TypicalWrapper) Wrap(c *Context) (err error) {

	if c.ProjectPackage == "" {
		c.ProjectPackage = retrieveProjectPackage()
	}

	// NOTE: create tmp folder if not exist
	os.MkdirAll(c.TmpFolder+"/build-tool", os.ModePerm)
	os.MkdirAll(c.TmpFolder+"/bin", os.ModePerm)

	cksmFile := c.TmpFolder + "/checksum"
	out := c.TmpFolder + "/bin/build-tool"
	srcPath := c.TmpFolder + "/build-tool/main.go"
	descriptorPkg := c.ProjectPackage + "/typical"

	var cksm *Checksum
	if cksm, err = CreateChecksum("typical"); err != nil {
		return
	}

	if _, err = os.Stat(out); os.IsNotExist(err) || !cksm.IsSame(cksmFile) {
		c.Info("Update checksum")
		if err = cksm.Save(cksmFile); err != nil {
			return
		}

		if _, err = os.Stat(srcPath); os.IsNotExist(err) {
			c.Infof("Generate build-tool main source: %s", srcPath)
			if err = typcore.WriteBuildToolMain(c.Ctx, srcPath, descriptorPkg); err != nil {
				return
			}
		}

		c.Info("Build the build-tool")
		return buildkit.NewGoBuild(out, srcPath).
			SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage", c.ProjectPackage).
			SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultTmpFolder", c.TmpFolder).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			WithStdin(os.Stdin).
			Execute(c.Ctx)
	}

	return
}

func retrieveProjectPackage() (pkg string) {
	var (
		err  error
		root string
		f    *os.File
	)

	if root, err = os.Getwd(); err != nil {
		panic(err.Error())
	}

	if f, err = os.Open(root + "/go.mod"); err != nil {
		// NOTE: go.mod is not exist. Check if the project sit in $GOPATH
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			return root[len(gopath):]
		}
		panic("Failed to retrieve ProjectPackage: `go.mod` is missing and the project not in $GOPATH")
	}
	defer f.Close()

	modfile := common.ParseModfile(f)
	return modfile.ProjectPackage
}
