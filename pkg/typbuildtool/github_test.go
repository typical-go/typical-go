package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestGithub(t *testing.T) {
	t.Run("SHOULD implement of ReleaseFilter", func(t *testing.T) {
		var _ typbuildtool.ReleaseFilter = typbuildtool.CreateGithub("owner", "repo")
	})
}
