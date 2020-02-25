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
	NoDelete bool
}

func (b *BuildTool) mock(ctx context.Context, c *typbuild.Context, opt *MockOption) (err error) {
	var (
		targets []*mockTarget
	)

	if !opt.NoDelete {
		os.RemoveAll(c.MockFolder)
		files, _ := filepath.Glob(c.MockFolder + "*")
		for _, f := range files {
			log.Infof("Remove %s", f)
			os.RemoveAll(f)
		}
	}

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
			"-destination", target.mockDest,
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
	srcPkg   string
	srcName  string
	mockPkg  string
	mockDest string
}

func createMockTarget(c *typbuild.Context, decl *prebld.Declaration) *mockTarget {
	var (
		mockPkg  = fmt.Sprintf("%s_%s", c.MockFolder, decl.File.Name.Name)
		mockDest = fmt.Sprintf("%s/%s.go", mockPkg, strcase.ToSnake(decl.SourceName))
		srcPkg   = fmt.Sprintf("%s/%s", c.ModulePackage, filepath.Dir(decl.Path))
	)
	return &mockTarget{
		srcPkg:   srcPkg,
		srcName:  decl.SourceName,
		mockPkg:  mockPkg,
		mockDest: mockDest,
	}
}
