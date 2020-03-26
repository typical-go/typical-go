package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestConstructor(t *testing.T) {
	t.Run("SHOULD implement Provider", func(t *testing.T) {
		var _ typapp.Provider = typapp.NewConstructor(nil)
	})
}
