package typmock_test

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
)

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

	directives := []*typast.Directive{
		{
			TagName: "@mock",
			Decl: &typast.Decl{
				File: typast.File{Package: "mypkg", Path: "parent/path/some_interface.go"},
				Type: &typast.InterfaceDecl{TypeDecl: typast.TypeDecl{Name: "SomeInterface"}},
			}},
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

	gomock := &typmock.GoMock{}
	require.NoError(t, gomock.Process(c, directives))
	require.Equal(t, "> rm -rf parent/path_mock\n> go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen\n> .typical-tmp/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface\n> Fail to mock '/parent/path.SomeInterface': some-error\n", out.String())
}
