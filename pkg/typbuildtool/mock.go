package typbuildtool

import (
	"go/build"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/coll"
	"github.com/urfave/cli"
)

func (t buildtool) cmdMock() cli.Command {
	return cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-delete",
				Usage: "Generate mock class with delete previous generation",
			},
		},
		Action: t.generateMock,
	}
}

func (t buildtool) generateMock(ctx *cli.Context) (err error) {
	log.Info("Generate mocks")
	if err = exec.Command("go", "get", "github.com/golang/mock/mockgen").Run(); err != nil {
		return
	}
	mockPkg := typenv.Layout.Mock
	if !ctx.Bool("no-delete") {
		log.Infof("Clean mock package '%s'", mockPkg)
		os.RemoveAll(mockPkg)
	}
	var errs coll.Errors
	for _, mockTarget := range t.MockTargets {
		dest := mockPkg + "/" + mockTarget[strings.LastIndex(mockTarget, "/")+1:]
		cmd := exec.Command(build.Default.GOPATH+"/bin/mockgen",
			"-source", mockTarget,
			"-destination", dest,
			"-package", mockPkg)
		if err := cmd.Run(); err != nil {
			errs.Append(err)
		}
	}
	return errs.ToError()
}
