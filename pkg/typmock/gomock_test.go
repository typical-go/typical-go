package typmock_test

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
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
			Annots: []*typast.Directive{
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

	GoMock := &typmock.GoMock{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "rm -rf parent/path_mock"},
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: "/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface"},
	})(t)
	require.NoError(t, GoMock.Execute(c))
}

func TestMockGen_InstallMockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp2"
	defer func() { typgo.TypicalTmp = "" }()

	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "go build -o .typical-tmp2/bin/mockgen github.com/golang/mock/mockgen", ReturnError: errors.New("some-error")},
	})(t)

	err := typmock.MockGen(c, "", "", "", "")
	require.EqualError(t, err, "some-error")
}

func TestAnnotate_MockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	summary := &typast.Summary{
		Annots: []*typast.Directive{
			{
				TagName: "@mock",
				Decl: &typast.Decl{
					File: typast.File{Package: "mypkg", Path: "parent/path/some_interface.go"},
					Type: &typast.InterfaceDecl{TypeDecl: typast.TypeDecl{Name: "SomeInterface"}},
				}},
		},
	}

	var out strings.Builder
	c := &typgo.Context{
		Context:    cli.NewContext(nil, &flag.FlagSet{}, nil),
		Descriptor: &typgo.Descriptor{},
		Logger:     typgo.Logger{Stdout: &out},
	}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "rm -rf parent/path_mock"},
		{CommandLine: "go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: ".typical-tmp/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface", ReturnError: errors.New("some-error")},
	})(t)

	require.NoError(t, typmock.Annotate(c, summary))
	require.Equal(t, "> rm -rf parent/path_mock\n> go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen\n> .typical-tmp/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface\n> Fail to mock '/parent/path.SomeInterface': some-error\n", out.String())
}
