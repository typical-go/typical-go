package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestConfiguration(t *testing.T) {
	cfg := typcore.NewConfiguration("some-name", "some-spec")
	require.Equal(t, "some-name", cfg.Name())
	require.Equal(t, "some-spec", cfg.Spec())
}
