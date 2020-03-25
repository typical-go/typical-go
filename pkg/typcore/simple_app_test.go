package typcore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestSimpleApp(t *testing.T) {
	t.Run("SHOULD implement App", func(t *testing.T) {
		var _ typcore.App = typcore.NewApp(nil)
	})

	t.Run("SHOULD implement AppSourceable", func(t *testing.T) {
		var _ typcore.AppSourceable = typcore.NewApp(nil)
	})

	t.Run("", func(t *testing.T) {
		fn := func(*typcore.Descriptor) error {
			return errors.New("some-error")
		}
		app := typcore.NewApp(fn, "some-source")
		require.EqualError(t, app.RunApp(nil), "some-error")
		require.Equal(t, []string{"some-source"}, app.AppSources())
	})
}
