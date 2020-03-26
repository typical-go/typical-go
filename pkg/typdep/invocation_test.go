package typdep_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typdep"
)

func TestInvocation(t *testing.T) {
	t.Run("SHOULD implement Invokable", func(t *testing.T) {
		var _ typdep.Invokable = typdep.NewInvocation(nil)
	})
}
