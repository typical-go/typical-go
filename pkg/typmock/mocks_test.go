package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestCreateMock(t *testing.T) {
	typgo.ProjectPkg = "some-proj"
	defer func() { typgo.ProjectPkg = "" }()

	testcases := []struct {
		testName string
		annot    *typgen.Directive
		expected *typmock.Mock
	}{
		{
			annot: &typgen.Directive{
				Decl: &typgen.Decl{
					File: typgen.File{
						Package: "somePkg",
						Path:    "path/folder/source.go",
					},
					Type: &typgen.InterfaceDecl{
						TypeDecl: typgen.TypeDecl{Name: "SomeInterface"},
					},
				},
				TagName: "mock",
			},
			expected: &typmock.Mock{
				Pkg:     "some-proj/path/folder",
				Source:  "SomeInterface",
				MockPkg: "somePkg_mock",
				Dest:    "internal/generated/mock/path/somePkg_mock/some_interface.go",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typmock.CreateMock(tt.annot))
		})
	}
}

func TestGenTarget(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	testcases := []struct {
		TestName string
		Dir      string
		Expected string
	}{
		{
			Dir:      "internal/app/service",
			Expected: "internal/generated/mock/app",
		},
		{
			Dir:      "internal/service",
			Expected: "internal/generated/mock",
		},
		{
			Dir:      ".",
			Expected: "internal/generated/mock",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typmock.GenTarget(tt.Dir))
		})
	}
}
