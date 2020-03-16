package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewApp(t *testing.T) {
	t.Run("SHOULD implement App", func(t *testing.T) {
		var _ typcore.App = typapp.New(nil)
	})
	t.Run("SHOULD implement Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = typapp.New(nil)
	})
	t.Run("SHOULD implement Provider", func(t *testing.T) {
		var _ typapp.Provider = typapp.New(nil)
	})
	t.Run("SHOULD implement Destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = typapp.New(nil)
	})
	t.Run("SHOULD implement Preparer", func(t *testing.T) {
		var _ typapp.Preparer = typapp.New(nil)
	})
	t.Run("SHOULD implement EntryPointer", func(t *testing.T) {
		var _ typapp.EntryPointer = typapp.New(nil)
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typapp.Commander = typapp.New(nil)
	})
	t.Run("SHOULD implement Sourceable", func(t *testing.T) {
		var _ typcore.SourceableApp = typapp.New(nil)
	})

}
