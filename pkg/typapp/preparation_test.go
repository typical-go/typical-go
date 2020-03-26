package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestPreparation(t *testing.T) {
	t.Run("SHOULD implement Preparer", func(t *testing.T) {
		var _ typapp.Preparer = typapp.NewPreparation(nil)
	})
}
