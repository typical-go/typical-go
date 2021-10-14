package typmock_test

import (
	"errors"
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
		annot    *typgen.Annotation
		expected *typmock.Mock
	}{
		{
			annot: &typgen.Annotation{
				Decl: &typgen.Decl{
					File: &typgen.File{Name: "folder", Path: "path/folder/source.go"},
					Type: &typgen.Interface{TypeDecl: typgen.TypeDecl{Name: "SomeInterface"}},
				},
				Name: "mock",
			},
			expected: &typmock.Mock{
				Package:            "some-proj/path/folder",
				Source:             "SomeInterface",
				DestinationPackage: "folder_mock",
				Destination:        "internal/generated/path/folder_mock/some_interface.go",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typmock.CreateMock(tt.annot))
		})
	}
}

func TestMock_Generate_WhenInstallMockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp2"
	defer func() { typgo.TypicalTmp = "" }()

	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go build -o .typical-tmp2/bin/mockgen github.com/golang/mock/mockgen", ReturnError: errors.New("some-error")},
	})(t)

	mock := &typmock.Mock{}
	err := mock.Generate(c)
	require.EqualError(t, err, "some-error")
}
