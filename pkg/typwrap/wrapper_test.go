package typwrap_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typwrap"
)

func TestNew(t *testing.T) {
	t.Run("SHOULD implement wrapper", func(t *testing.T) {
		var _ typwrap.Wrapper = typwrap.New()
	})
}
