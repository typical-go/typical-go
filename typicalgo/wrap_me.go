package typicalgo

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
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
	"github.com/urfave/cli/v2"
)

type wrapContext struct {
	*typcore.Context
	Cli            *cli.Context
	tmp            string
	projectPackage string
}

func wrapMe(c *wrapContext) (err error) {

	if c.projectPackage == "" {
		c.projectPackage = retrieveProjectPackage(c)
	}

	// NOTE: create tmp folder if not exist
	typcore.MakeTempDir(c.tmp)

	checksumPath := typcore.Checksum(c.tmp)
	buildToolBin := typcore.BuildToolBin(c.tmp)
	var checksumData []byte
	if checksumData, err = checksum("typical"); err != nil {
		return
	}

	if !sameChecksum(checksumPath, checksumData) || notExist(buildToolBin) {
		c.Info("Update new checksum")
		if err = ioutil.WriteFile(checksumPath, checksumData, 0777); err != nil {
			return
		}
		c.Info("Build the Build-Tool")
		if err = buildBuildTool(c); err != nil {
			return
		}
	}

	return
}

func buildBuildTool(c *wrapContext) (err error) {
	var (
		descriptorPkg = typcore.TypicalPackage(c.projectPackage)
		srcPath       = typcore.BuildToolSrc(c.tmp)
		binPath       = typcore.BuildToolBin(c.tmp)
		ctx           = c.Cli.Context
	)

	if notExist(srcPath) {
		c.Infof("Generate Build-Tool main source: %s", srcPath)
		if err = exor.NewWriteTemplate(srcPath, tmpl.MainSrcBuildTool, tmpl.MainSrcData{
			DescriptorPackage: descriptorPkg,
		}).Execute(ctx); err != nil {
			return
		}
	}

	return exor.NewGoBuild(binPath, srcPath).
		SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage", c.projectPackage).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithStdin(os.Stdin).
		Execute(ctx)
}

func notExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func checksum(path string) ([]byte, error) {
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

func retrieveProjectPackage(c *wrapContext) (pkg string) {
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
