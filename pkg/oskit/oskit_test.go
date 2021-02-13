package oskit_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
)

func TestMkdirAll(t *testing.T) {
	func() {
		defer oskit.MkdirAll("some-dir")()
		_, err := os.Stat("some-dir")
		require.False(t, os.IsNotExist(err))
	}()
	_, err := os.Stat("some-dir")
	require.True(t, os.IsNotExist(err))
}
