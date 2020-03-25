package typcfg_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcfg"

	"github.com/stretchr/testify/require"
)

func TestConfiguration(t *testing.T) {
	t.Run("SHOULD implement typcfg.Configurer", func(t *testing.T) {
		var _ typcfg.Configurer = typcfg.NewConfiguration("some-name", "some-spec")
	})

	t.Run("", func(t *testing.T) {
		cfg := typcfg.NewConfiguration("some-name", "some-spec")
		require.Equal(t, "some-name", cfg.Name)
		require.Equal(t, "some-spec", cfg.Spec)
	})
}
