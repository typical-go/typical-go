package typbuild

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
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (*Build) cmdMock(c *Context) *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
		},
		Action: func(cliCtx *cli.Context) (err error) {
			var (
				targets common.Strings
				mockgen string
				errs    common.Errors
				ctx     = cliCtx.Context
			)
			if err = c.EachAnnotation("mock", walker.InterfaceType, func(decl *walker.Declaration, ann *walker.Annotation) (err error) {
				targets.Append(decl.Filename)
				return
			}); err != nil {
				return
			}
			if mockgen, err = installMockgen(ctx); err != nil {
				return
			}
			mockPkg := typenv.Layout.Mock
			if !cliCtx.Bool("no-delete") {
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
