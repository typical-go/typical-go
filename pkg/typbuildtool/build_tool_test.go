package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestBuildTool(t *testing.T) {
	t.Run("SHOULD implement Buildtool", func(t *testing.T) {
		var _ typcore.BuildTool = typbuildtool.Create()
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typbuildtool.Commander = typbuildtool.Create()
	})
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.Create()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Tester = typbuildtool.Create()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.Create()
	})
	t.Run("SHOULD implement Releaser", func(t *testing.T) {
		var _ typbuildtool.Releaser = typbuildtool.Create()
	})
	t.Run("SHOULD implement Publisher", func(t *testing.T) {
		var _ typbuildtool.Publisher = typbuildtool.Create()
	})
	t.Run("SHOULD implement Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = typbuildtool.Create()
	})
}
