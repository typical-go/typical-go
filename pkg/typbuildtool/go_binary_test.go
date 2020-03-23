package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestBuildDistribution(t *testing.T) {
	t.Run("SHOULD implement BuildDistribution", func(t *testing.T) {
		var _ typbuildtool.BuildDistribution = &typbuildtool.GoBinary{}
	})
}
