package typbuildtool

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iancoleman/strcase"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
)

// MockOption is option for generate mock
type MockOption struct {
}

func (b *BuildTool) mock(ctx context.Context, c *typbuild.Context, opt *MockOption) (err error) {
	var (
		targets []*mockTarget
	)

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.TempFolder)

	if err = c.EachAnnotation("mock", prebld.InterfaceType, func(decl *prebld.Declaration, ann *prebld.Annotation) (err error) {
		targets = append(targets, createMockTarget(c, decl))
		return
	}); err != nil {
		return
	}

	if !common.IsFileExist(mockgen) {
		log.Info("Build mockgen")
		if err = buildkit.NewGoBuild(mockgen, "github.com/golang/mock/mockgen").Command(ctx).Run(); err != nil {
			return
		}
	}

	for _, target := range targets {
		log.Infof("Mock %s", target.srcName)
		cmd := exec.CommandContext(ctx, mockgen,
			"-destination", target.dest,
			"-package", target.mockPkg,
			target.srcPkg,
			target.srcName,
		)
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Errorf("Mock '%s' failed: %w", target, err)
		}
	}
	return
}

type mockTarget struct {
	srcPkg  string
	srcName string
	mockPkg string
	dest    string
}

func createMockTarget(c *typbuild.Context, decl *prebld.Declaration) *mockTarget {
	var (
		pkg     = decl.File.Name.Name
		dir     = filepath.Dir(decl.Path)
		dirDest = dir[:len(dir)-len(pkg)]
		srcPkg  = fmt.Sprintf("%s/%s", c.ModulePackage, dir)
		mockPkg = fmt.Sprintf("%s_%s", c.MockFolder, pkg)
		dest    = fmt.Sprintf("%s%s/%s.go", dirDest, mockPkg, strcase.ToSnake(decl.SourceName))
	)
	return &mockTarget{
		srcPkg:  srcPkg,
		srcName: decl.SourceName,
		mockPkg: mockPkg,
		dest:    dest,
	}
}
