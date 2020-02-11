package app_test

import (
	"testing"

	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement AppCommander", func(t *testing.T) {
		var _ typapp.AppCommander = &app.Module{}
	})

}
