package typmock_test

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

func TestExecute(t *testing.T) {
	defer monkey.Patch(filekit.FindDir, func(includes, excludes []string) ([]string, error) {
		return nil, nil
	}).Unpatch()

	defer monkey.Patch(typast.Walk, func(layouts []string) (dirs, files []string) {
		return
	}).Unpatch()

	defer monkey.Patch(typast.Compile, func(paths ...string) (*typast.Summary, error) {
		return &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@mock",
					Decl: &typast.Decl{
						File: typast.File{Package: "mypkg", Path: "parent/path/some_interface.go"},
						Type: &typast.InterfaceDecl{TypeDecl: typast.TypeDecl{Name: "SomeInterface"}},
					},
				},
			},
		}, nil
	}).Unpatch()

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "rm -rf parent/path_mock"},
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: "/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface"},
	})(t)

	generateMock := &typmock.GenerateMock{}
	c, _ := typgo.DummyContext()
	require.NoError(t, generateMock.Execute(c))
}

func TestMockGen_InstallMockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp2"
	defer func() { typgo.TypicalTmp = "" }()

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -o .typical-tmp2/bin/mockgen github.com/golang/mock/mockgen", ReturnError: errors.New("some-error")},
	})(t)

	c, _ := typgo.DummyContext()
	err := typmock.MockGen(c, "", "", "", "")
	require.EqualError(t, err, "some-error")
}

func TestAnnotate_MockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	summary := &typast.Summary{
		Annots: []*typast.Annot{
			{
				TagName: "@mock",
				Decl: &typast.Decl{
					File: typast.File{Package: "mypkg", Path: "parent/path/some_interface.go"},
					Type: &typast.InterfaceDecl{TypeDecl: typast.TypeDecl{Name: "SomeInterface"}},
				}},
		},
	}

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "rm -rf parent/path_mock"},
		{CommandLine: "go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: ".typical-tmp/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface", ReturnError: errors.New("some-error")},
	})(t)

	c, out := typgo.DummyContext()
	require.NoError(t, typmock.Annotate(c, summary))
	require.Equal(t, "some-project:dummy> $ rm -rf parent/path_mock\nsome-project:dummy> $ .typical-tmp/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface\nsome-project:dummy> Fail to mock '/parent/path.SomeInterface': some-error\n", out.String())
}
