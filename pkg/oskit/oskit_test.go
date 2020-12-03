package oskit_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
)

func TestPatchStdout(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()
	fmt.Fprintln(oskit.Stdout, "some-text")
	require.Equal(t, "some-text\n", out.String())
}

func TestMkdirAll(t *testing.T) {
	func() {
		defer oskit.MkdirAll("some-dir")()
		_, err := os.Stat("some-dir")
		require.False(t, os.IsNotExist(err))
	}()
	_, err := os.Stat("some-dir")
	require.True(t, os.IsNotExist(err))
}
