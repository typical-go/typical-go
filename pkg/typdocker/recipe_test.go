package typdocker_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typdocker"
)

func TestRecipe(t *testing.T) {
	t.Run("SHOULD implement composer", func(t *testing.T) {
		var _ typdocker.Composer = &typdocker.Recipe{}
	})
}
