package typbuildtool

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func cmdMock(d *typcore.ProjectDescriptor) *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
		},
		Action: func(c *cli.Context) (err error) {
			var (
				targets  Automocks
				mockgen  string
				errs     common.Errors
				projInfo typcore.ProjectInfo
				events   walker.Declarations
				ctx      = c.Context
			)
			if projInfo, err = typcore.ReadProject(typenv.Layout.App); err != nil {
				log.Fatal(err.Error())
			}
			log.Info("Walk the project")
			if events, err = walker.Walk(projInfo.Files); err != nil {
				return
			}
			if err = events.EachAnnotation("mock", walker.InterfaceType, targets.OnAnnotation); err != nil {
				return
			}
			targets = append(targets, d.MockTargets...)
			if mockgen, err = installMockgen(ctx); err != nil {
				return
			}
			mockPkg := typenv.Layout.Mock
			if !c.Bool("no-delete") {
				log.Infof("Clean mock package '%s'", mockPkg)
				os.RemoveAll(mockPkg)
			}
			for _, target := range targets {
				log.Infof("Mock '%s'", target)
				dest := mockPkg + "/" + filepath.Base(target)
				cmd := exec.CommandContext(ctx,
					mockgen,
					"-source", target,
					"-destination", dest,
					"-package", mockPkg,
				)
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					errs.Append(fmt.Errorf("Mock '%s' failed: %w", target, err))
				}
			}
			return errs.Unwrap()
		},
	}
}

func installMockgen(ctx context.Context) (path string, err error) {
	path = build.Default.GOPATH + "/bin/mockgen"
	err = exec.CommandContext(ctx, "go", "get", "github.com/golang/mock/mockgen").Run()
	return
}
