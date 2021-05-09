package typgen_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestLayouts(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	walker := typgen.Layouts{"pkg", "wrapper"}

	require.Equal(t, []string{
		"pkg/some_lib/lib.go",
		"wrapper/some_pkg/some_file.go",
	}, walker.Walk())
}

func TestFilePaths(t *testing.T) {
	walker := typgen.FilePaths{"1", "2"}
	require.Equal(t, []string{"1", "2"}, walker.Walk())
}
