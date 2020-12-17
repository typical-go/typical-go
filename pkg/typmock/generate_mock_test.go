package typmock_test

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"github.com/typical-go/typical-go/pkg/oskit"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
)

func TestCommand(t *testing.T) {
	mockCmd := &typmock.GenerateMock{}
	command := mockCmd.Task(&typgo.BuildSys{})

	require.Equal(t, "mock", command.Name)
}

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
	require.NoError(t, generateMock.Execute(&typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}))
}

func TestMockGen_InstallMockgenError(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen", ReturnError: errors.New("some-error")},
	})(t)

	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	err := typmock.MockGen(c, "", "", "", "")
	require.EqualError(t, err, "some-error")
}

func TestAnnotate_MockgenError(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()

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
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: "/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface", ReturnError: errors.New("some-error")},
	})(t)

	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	require.NoError(t, typmock.Annotate(c, summary))
	require.Equal(t, "\n$ rm -rf parent/path_mock\n\n$ go build -o /bin/mockgen github.com/golang/mock/mockgen\n\n$ /bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface\nFail to mock '/parent/path.SomeInterface': some-error\n", out.String())
}
