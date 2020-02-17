package typicalgo_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/sample/app"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement AppCommander", func(t *testing.T) {
		var _ typapp.AppCommander = app.New()
	})

}
