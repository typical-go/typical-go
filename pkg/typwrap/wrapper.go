package typwrap

import (
	"bufio"
	"errors"
	"go/build"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
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
		if c.ProjectPackage, err = retrieveProjectPackage(); err != nil {
			return
		}
	}

	// NOTE: create tmp folder if not exist
	os.MkdirAll(c.TypicalTmp+"/build-tool", os.ModePerm)
	os.MkdirAll(c.TypicalTmp+"/bin", os.ModePerm)

	cksmFile := c.TypicalTmp + "/checksum"
	out := c.TypicalTmp + "/bin/build-tool"
	srcPath := c.TypicalTmp + "/build-tool/main.go"
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
			SetVariable(typcore.DefaultProjectPackageVar, c.ProjectPackage).
			SetVariable(typcore.DefaultTypicalTmpVar, c.TypicalTmp).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			WithStdin(os.Stdin).
			Execute(c.Ctx)
	}

	return
}

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
