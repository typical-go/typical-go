package typmock

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"
)

// StdMocker is standard mocker
type StdMocker struct {
}

// New return new instance of StdMocker
func New() *StdMocker {
	return &StdMocker{}
}

// Mock the project
func (b *StdMocker) Mock(ctx context.Context, c *Context) (err error) {
	var (
		targets []*mockTarget
	)

	if err = c.EachAnnotation("mock", typast.InterfaceType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		targets = append(targets, createMockTarget(c, decl))
		return
	}); err != nil {
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.TempFolder)

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
