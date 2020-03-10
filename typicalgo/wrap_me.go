package typicalgo

import (
	"bytes"
	"context"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/buildkit"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
)

type wrapContext struct {
	*typcore.Descriptor
	tmp            string
	projectPackage string
}

func wrapMe(ctx context.Context, wc *wrapContext) (err error) {

	// NOTE: create tmp folder if not exist
	typcore.MakeTempDir(wc.tmp)

	checksumPath := typcore.Checksum(wc.tmp)
	buildToolBin := typcore.BuildToolBin(wc.tmp)
	var checksumData []byte
	if checksumData, err = checksum("typical"); err != nil {
		return
	}

	if !sameChecksum(checksumPath, checksumData) || notExist(buildToolBin) {
		log.Info("Update new checksum")
		if err = ioutil.WriteFile(checksumPath, checksumData, 0777); err != nil {
			return
		}
		log.Info("Build the Build-Tool")
		if err = buildBuildTool(ctx, wc); err != nil {
			return
		}
	}

	return
}

func buildBuildTool(ctx context.Context, wc *wrapContext) (err error) {
	var (
		descriptorPkg = typcore.TypicalPackage(wc.projectPackage)
		srcPath       = typcore.BuildToolSrc(wc.tmp)
		binPath       = typcore.BuildToolBin(wc.tmp)
	)

	if notExist(srcPath) {
		data := tmpl.MainSrcData{
			DescriptorPackage: descriptorPkg,
		}
		if err = runnerkit.NewWriteTemplate(srcPath, tmpl.MainSrcBuildTool, data).Run(ctx); err != nil {
			return
		}
	}

	gobuild := buildkit.NewGoBuild(binPath, srcPath)
	gobuild.SetVariable("github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage", wc.projectPackage)

	cmd := gobuild.Command(ctx)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
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
