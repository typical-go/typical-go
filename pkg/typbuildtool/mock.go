package typbuildtool

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdMock() *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
		},
		Action: t.generateMock,
	}
}

func (t buildtool) generateMock(ctx *cli.Context) (err error) {
	var targets Automocks
	walker := walker.New(t.filenames)
	walker.AddTypeSpecListener(&targets)
	log.Info("Walk the project")
	if err = walker.Walk(); err != nil {
		return
	}
	targets = append(targets, t.MockTargets...)
	var mockgen string
	if mockgen, err = installMockgen(); err != nil {
		return
	}
	mockPkg := typenv.Layout.Mock
	if !ctx.Bool("no-delete") {
		log.Infof("Clean mock package '%s'", mockPkg)
		os.RemoveAll(mockPkg)
	}
	var errs common.Errors
	for _, target := range targets {
		log.Infof("Mock '%s'", target)
		dest := mockPkg + "/" + filepath.Base(target)
		cmd := exec.Command(mockgen,
			"-source", target,
			"-destination", dest,
			"-package", mockPkg)
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			errs.Append(fmt.Errorf("Mock '%s' failed: %w", target, err))
		}
	}
	return errs.Unwrap()
}

func installMockgen() (path string, err error) {
	path = build.Default.GOPATH + "/bin/mockgen"
	err = exec.Command("go", "get", "github.com/golang/mock/mockgen").Run()
	return
}
