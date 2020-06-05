package typmock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestCreateMock(t *testing.T) {
	testcases := []struct {
		testName string
		annot    *typast.Annot
		expected *typmock.Mock
	}{
		{
			annot: &typast.Annot{
				Decl: &typast.Decl{
					Pkg:  "somePkg",
					Path: "/path/folder/source.go",
					Type: typast.Interface,
					Name: "SomeInterface",
				},
				TagName: "mock",
			},
			expected: &typmock.Mock{
				Dir:     "/path/folder",
				Pkg:     "somePkg",
				Source:  "SomeInterface",
				Parent:  "/path",
				MockPkg: "somePkg_mock",
			},
		},
		{
			annot: &typast.Annot{
				Decl: &typast.Decl{
					Pkg:  "somePkg",
					Path: "/path/folder/source.go",
					Type: typast.Function,
					Name: "SomeInterface",
				},
				TagName: "mock",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typmock.CreateMock(tt.annot))
		})
	}

}
