package typbuildtool

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
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
	if err = bash.GoGet("github.com/golang/mock/mockgen"); err != nil {
		return
	}
	mockPkg := typenv.Mock
	if !ctx.Bool("no-delete") {
		log.Infof("Clean mock package '%s'", mockPkg)
		os.RemoveAll(mockPkg)
	}
	for _, mockTarget := range t.MockTargets {
		dest := mockPkg + "/" + mockTarget[strings.LastIndex(mockTarget, "/")+1:]
		err = bash.RunGoBin("mockgen",
			"-source", mockTarget,
			"-destination", dest,
			"-package", mockPkg)
	}
	return
}
