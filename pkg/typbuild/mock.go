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
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// MockOption is option for generate mock
type MockOption struct {
	NoDelete bool
}

func (b *Build) mock(ctx context.Context, c *Context, opt *MockOption) (err error) {
	var (
		targets common.Strings
		mockgen string
		errs    common.Errors
		mockPkg = typenv.Layout.Mock
	)

	if !opt.NoDelete {
		log.Infof("Clean mock package '%s'", mockPkg)
		os.RemoveAll(mockPkg)
	}
	if err = c.EachAnnotation("mock", prebld.InterfaceType, func(decl *prebld.Declaration, ann *prebld.Annotation) (err error) {
		targets.Append(decl.Filename)
		return
	}); err != nil {
		return
	}
	if mockgen, err = installMockgen(ctx); err != nil {
		return
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
}

func installMockgen(ctx context.Context) (path string, err error) {
	path = build.Default.GOPATH + "/bin/mockgen"
	err = exec.CommandContext(ctx, "go", "get", "github.com/golang/mock/mockgen").Run()
	return
}
