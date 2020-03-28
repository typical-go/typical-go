package typdocker_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

func TestDocker(t *testing.T) {
	t.Run("SHOULD implement Task", func(t *testing.T) {
		var _ typbuildtool.Utility = typdocker.Compose()
	})
}
