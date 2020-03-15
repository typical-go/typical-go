package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestStdBuilder(t *testing.T) {
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.NewBuilder()
	})
	t.Run("SHOULD implement Cleaner", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.NewBuilder()
	})
}
