package typbuildtool

import (
	"bytes"
	"crypto/sha256"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Wrap the project
func (b *TypicalBuildTool) Wrap(c *typcore.WrapContext) (err error) {

	if c.ProjectPackage == "" {
		c.ProjectPackage = retrieveProjectPackage(c)
	}

	// NOTE: create tmp folder if not exist
	os.MkdirAll(c.TmpFolder+"/build-tool", os.ModePerm)
	os.MkdirAll(c.TmpFolder+"/bin", os.ModePerm)

	checksum := c.TmpFolder + "/checksum"
	out := c.TmpFolder + "/bin/build-tool"
	srcPath := c.TmpFolder + "/build-tool/main.go"

	var checksumData []byte
	if checksumData, err = generateChecksum("typical"); err != nil {
		return
	}

	if !sameChecksum(checksum, checksumData) || notExist(out) {
		var (
			descriptorPkg = c.ProjectPackage + "/typical"
		)

		c.Info("Update new checksum")
		if err = ioutil.WriteFile(checksum, checksumData, 0777); err != nil {
			return
		}

		if notExist(srcPath) {
			c.Infof("Generate Build-Tool main source: %s", srcPath)
			if err = exor.NewWriteTemplate(srcPath, tmpl.MainSrcBuildTool, tmpl.MainSrcData{
				DescriptorPackage: descriptorPkg,
			}).Execute(c.Ctx); err != nil {
				return
			}
		}

		c.Info("Build the Build-Tool")
		return exor.NewGoBuild(out, srcPath).
			SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage", c.ProjectPackage).
			SetVariable("github.com/typical-go/typical-go/pkg/typbuildtool.DefaultTmpFolder", c.TmpFolder).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			WithStdin(os.Stdin).
			Execute(c.Ctx)
	}

	return
}

func notExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func generateChecksum(path string) ([]byte, error) {
	h := sha256.New()
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(h, path)
		return nil
	}); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func sameChecksum(path string, data []byte) bool {
	var (
		current []byte
		err     error
	)
	if current, err = ioutil.ReadFile(path); err != nil {
		return false
	}
	return bytes.Compare(current, data) == 0
}

func retrieveProjectPackage(c *typcore.WrapContext) (pkg string) {
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
