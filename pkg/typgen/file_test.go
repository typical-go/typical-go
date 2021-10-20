package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestFile_SourceCode(t *testing.T) {
	testCases := []struct {
		TestName string
		File     *typgen.File
		Expected string
	}{
		{
			File: &typgen.File{
				Name: "some package",
				Import: []*typgen.Import{
					{Name: "", Path: "fmt"},
					{Name: "a", Path: "github.com/typical-go/typical-go"},
				},
			},
			Expected: `package some package

import (
	"fmt"
	a "github.com/typical-go/typical-go"
)`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.File.Code())
		})
	}
}

func TestPackageName(t *testing.T) {
	testcases := []struct {
		TestName string
		Path     string
		Expected string
	}{
		{
			Path:     "a/b/c/file.go",
			Expected: "c",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typgen.PackageName(tt.Path))
		})
	}
}
