package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestBuildTool(t *testing.T) {
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typbuildtool.Commander = typbuildtool.New()
	})
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.New()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Tester = typbuildtool.New()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.New()
	})
}
