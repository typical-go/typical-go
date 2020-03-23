package app_test

import (
	"testing"

	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typwrap"
)

func TestNew(t *testing.T) {
	t.Run("SHOULD implement app", func(t *testing.T) {
		var _ typcore.App = app.New()
	})
	t.Run("SHOULD implement wrapper", func(t *testing.T) {
		var _ typwrap.Wrapper = app.New()
	})
}
