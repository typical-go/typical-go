package buildkit_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGoMod(t *testing.T) {
	t.Run("", func(t *testing.T) {
		path := "go.mod"
		ioutil.WriteFile(path, []byte("module github.com/typical-go/typical-go\ngo 1.13"), 0644)
		defer os.Remove(path)

		gomod, err := buildkit.CreateGoMod(path)
		require.NoError(t, err)
		require.Equal(t, &buildkit.GoMod{
			ModulePackage: "github.com/typical-go/typical-go",
			GoVersion:     "1.13",
		}, gomod)
	})

	t.Run("WHEN path not exist", func(t *testing.T) {
		_, err := buildkit.CreateGoMod("not-exist")
		require.True(t, os.IsNotExist(err))
		require.EqualError(t, err, "open not-exist: no such file or directory")
	})

}
