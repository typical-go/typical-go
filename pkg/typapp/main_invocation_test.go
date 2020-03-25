package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestMainInvocation(t *testing.T) {
	t.Run("SHOULD implement EntryPointer", func(t *testing.T) {
		var _ typapp.EntryPointer = typapp.NewMainInvocation(nil)
	})
}
