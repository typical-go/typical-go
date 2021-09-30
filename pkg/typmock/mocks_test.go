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
					File: &typgen.File{Name: "folder", Path: "path/folder/source.go"},
					Type: &typgen.Interface{TypeDecl: typgen.TypeDecl{Name: "SomeInterface"}},
				},
				TagName: "mock",
			},
			expected: &typmock.Mock{
				Pkg:     "some-proj/path/folder",
				Source:  "SomeInterface",
				MockPkg: "folder_mock",
				Dest:    "internal/generated/path/folder_mock/some_interface.go",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typmock.CreateMock(tt.annot))
		})
	}
}

func TestGeneratedDir(t *testing.T) {
	testCases := []struct {
		TestName  string
		Directive *typgen.Directive
		Suffix    string
		Expected  string
	}{
		{
			Directive: &typgen.Directive{
				Decl: &typgen.Decl{
					File: &typgen.File{
						Path: ".",
					},
				},
			},
			Expected: "internal/generated",
		},
		{
			Directive: &typgen.Directive{
				Decl: &typgen.Decl{
					File: &typgen.File{
						Path: "internal/app/service/file.go",
					},
				},
			},
			Suffix:   "mock",
			Expected: "internal/generated/app/service_mock",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typmock.GeneratedDir(tt.Directive, tt.Suffix))
		})
	}
}
