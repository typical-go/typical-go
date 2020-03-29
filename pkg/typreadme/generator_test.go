package typreadme_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typreadme"
)

func TestReadme(t *testing.T) {
	t.Run("SHOULD implement of Task", func(t *testing.T) {
		var _ typbuildtool.Utility = typreadme.Generator()
	})
}
