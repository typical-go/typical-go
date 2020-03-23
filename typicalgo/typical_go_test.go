package typicalgo_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typwrap"
	"github.com/typical-go/typical-go/typicalgo"
)

func TestNew(t *testing.T) {
	t.Run("SHOULD implement app", func(t *testing.T) {
		var _ typcore.App = typicalgo.New()
	})
	t.Run("SHOULD implement wrapper", func(t *testing.T) {
		var _ typwrap.Wrapper = typicalgo.New()
	})
}
