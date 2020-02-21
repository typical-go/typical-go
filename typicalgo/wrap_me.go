package typicalgo

import (
	"bytes"
	"context"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/common/stdrun"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
)

func wrapMe(ctx context.Context, d *typcore.Descriptor, typTmp string) (err error) {

	var (
		descriptorPkg   = d.ModulePackage + "/typical"
		buildToolTarget = typTmp + "/bin/build-tool"
		buildToolMain   = typTmp + "/build-tool/main.go"
		checksumPath    = typTmp + "/checksum"
	)

	// NOTE: create typical tmp if not exist
	os.MkdirAll(typTmp+"/build-tool", os.ModePerm)
	os.MkdirAll(typTmp+"/bin", os.ModePerm)

	var checksumData []byte
	if checksumData, err = checksum("typical"); err != nil {
		return
	}

	if !sameChecksum(checksumPath, checksumData) || notExist(buildToolTarget) {
		log.Info("Update new checksum")
		if err = ioutil.WriteFile(checksumPath, checksumData, 0644); err != nil {
			return
		}
		log.Info("Build the Build-Tool")
		if err = buildBuildTool(ctx, buildToolMain, buildToolTarget, descriptorPkg); err != nil {
			return
		}
	}

	return
}

func buildBuildTool(ctx context.Context, mainPath, target, descriptorPkg string) (err error) {
	if notExist(mainPath) {
		data := tmpl.MainSrcData{
			DescriptorPackage: descriptorPkg,
		}
		stdrun.NewWriteTemplate(mainPath, tmpl.MainSrcBuildTool, data).Run()
	}

	cmd := exec.CommandContext(ctx, "go", "build", "-o", target, mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
