package typmock

import (
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
	targetMap map[string][]*Target
}

// New return new instance of StdMocker
func New() *StdMocker {
	return &StdMocker{
		targetMap: make(map[string][]*Target),
	}
}

// Put new target
func (b *StdMocker) Put(target *Target) {
	key := target.MockDir
	if _, ok := b.targetMap[key]; ok {
		b.targetMap[key] = append(b.targetMap[key], target)
	} else {
		b.targetMap[key] = []*Target{target}
	}
}

// TargetMap return targetMap field
func (b *StdMocker) TargetMap() map[string][]*Target {
	return b.targetMap
}

// Mock the project
func (b *StdMocker) Mock(c *Context) (err error) {
	if err = c.EachAnnotation("mock", typast.InterfaceType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		b.Put(createTarget(c, decl))
		return
	}); err != nil {
		return
	}

	mockgen := fmt.Sprintf("%s/bin/mockgen", c.TempFolder)
	ctx := c.Cli.Context

	if !common.IsFileExist(mockgen) {
		log.Info("Build mockgen")
		if err = buildkit.NewGoBuild(mockgen, "github.com/golang/mock/mockgen").Command(ctx).Run(); err != nil {
			return
		}
	}

	for pkg, targets := range b.targetMap {

		log.Infof("Remove package: %s", pkg)
		os.RemoveAll(pkg)

		for _, target := range targets {
			log.Infof("Generate mock: %s", target.Dest)
			cmd := exec.CommandContext(ctx, mockgen,
				"-destination", target.Dest,
				"-package", target.MockPkg,
				target.SrcPkg,
				target.SrcName,
			)
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Errorf("Mock '%s' failed: %w", target, err)
			}
		}
	}

	return
}
