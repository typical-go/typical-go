package typcore_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestWalkLayout(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	dirs, files := typcore.WalkLayout([]string{
		"pkg",
		"wrapper",
	})

	require.Equal(t, []string{
		"pkg",
		"pkg/some_lib",
		"wrapper",
		"wrapper/some_pkg",
	}, dirs)

	require.Equal(t, []string{
		"pkg/some_lib/lib.go",
		"wrapper/some_pkg/some_file.go",
	}, files)
}
