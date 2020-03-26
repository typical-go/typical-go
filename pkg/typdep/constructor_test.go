package typdep_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typdep"
)

func TestConstructor(t *testing.T) {
	t.Run("SHOULD implement Providable", func(t *testing.T) {
		var _ typdep.Providable = typdep.NewConstructor(nil)
	})
}
