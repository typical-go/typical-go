package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestDestruction(t *testing.T) {
	t.Run("SHOULD implement destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = typapp.NewDestruction(nil)
	})
}
